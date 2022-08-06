package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/smithy-go"
	db "github.com/danielmachado86/contracts/db"
	"github.com/danielmachado86/contracts/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Username  string `json:"username" binding:"required,alphanum"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	ChangedAt time.Time `json:"changedAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		ChangedAt: user.ChangedAt,
		CreatedAt: user.CreatedAt,
	}
}

func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		server.Logger.Errorf("failed to unmarshal createUser request body")
		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	passwordHashed, err := utils.HashPasword(req.Password)
	if err != nil {
		server.Logger.Errorf("failed to hash password")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	arg := db.CreateUserParams{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Username:       req.Username,
		Email:          req.Email,
		PasswordHashed: passwordHashed,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				server.Logger.Errorf("failure creating user caused by unique constraint violation")
				ctx.JSON(http.StatusConflict, errorResponse(err, http.StatusConflict))
				return
			}
		}
		var oe *smithy.OperationError
		if errors.As(err, &oe) {
			var re *awshttp.ResponseError
			if errors.As(oe.Err, &re) {
				var tc *types.TransactionCanceledException
				if errors.As(re.Err, &tc) {
					server.Logger.Errorf("error due to transaction cancelation")
					var rList []string
					for _, r := range tc.CancellationReasons {
						av := &db.AV{}
						if r.Message != nil {
							err := attributevalue.UnmarshalMap(r.Item, av)
							rList = append(rList, av.Pk)
							if err != nil {
								ctx.JSON(http.StatusBadRequest, errorResponse(
									errors.New("couldn't unmarshall dynamodb attributes"),
									http.StatusBadRequest,
								))
							}
						}
					}
					msg := fmt.Sprintf("failed to create user: %s fields already exists", strings.Join(rList, `, `))
					server.Logger.Error(msg)
					ctx.JSON(http.StatusBadRequest, errorResponse(
						errors.New(msg),
						http.StatusBadRequest,
					))
					return
				}
			}

		}

	}

	server.Logger.Infof("user %s succesfully created", req.Username)

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusCreated, rsp)
}

type CreateSessionRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type CreateSessionResponse struct {
	SessionId             string       `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

func (server *Server) createSessions(ctx *gin.Context) {
	var req CreateSessionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		server.Logger.Errorf("failed to unmarshal loginUser request body")
		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			server.Logger.Errorf("user %s not found", req.Username)
			ctx.JSON(http.StatusNotFound, errorResponse(err, http.StatusNotFound))
			return
		}
		if err == db.ErrNotFound {
			server.Logger.Errorf("user %s not found", req.Username)
			ctx.JSON(http.StatusNotFound, errorResponse(err, http.StatusNotFound))
			return
		}
		server.Logger.Errorf("failed to check if user exists")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	err = utils.CheckPassword(req.Password, user.PasswordHashed)
	if err != nil {
		server.Logger.Errorf("user %s not authorized", req.Username)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err, http.StatusUnauthorized))
		return
	}

	accessToken, accessTokenPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		server.Logger.Errorf("failed to create token", req.Username)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
	}

	refreshToken, refreshTokenPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		server.Logger.Errorf("failed to create token for user: ", req.Username)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshTokenPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshTokenPayload.ExpiredAt,
		CreatedAt:    refreshTokenPayload.IssuedAt,
	})
	if err != nil {
		server.Logger.Errorf("failed to create session for user: ", req.Username)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
	}

	rsp := CreateSessionResponse{
		SessionId:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessTokenPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshTokenPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}

	server.Logger.Infof("user %s succesfully authenticated", req.Username)
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) deleteSessions(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
