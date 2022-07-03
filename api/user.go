package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/danielmachado86/contracts/db/sqlc"
	"github.com/danielmachado86/contracts/utils"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	Name              string    `json:"name"`
	LastName          string    `json:"lastName"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	CreatedAt         time.Time `json:"createdAt"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Name:              user.Name,
		LastName:          user.LastName,
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
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

	hashedPassword, err := utils.HashPasword(req.Password)
	if err != nil {
		server.Logger.Errorf("failed to hash password")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	arg := db.CreateUserParams{
		Name:           req.Name,
		LastName:       req.LastName,
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: hashedPassword,
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
		server.Logger.Errorf("failed to create user")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	server.Logger.Infof("user %s succesfully created", req.Username)

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusCreated, rsp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) createSessions(ctx *gin.Context) {
	var req loginUserRequest
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
		server.Logger.Errorf("failed to check if user exists")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		server.Logger.Errorf("user %s not authorized", req.Username)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err, http.StatusUnauthorized))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		server.Logger.Errorf("failed to create token", req.Username)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	server.Logger.Infof("user %s succesfully authenticated", req.Username)
	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) deleteSessions(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}
