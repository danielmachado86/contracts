package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/danielmachado86/contracts/db/sqlc"
	"github.com/danielmachado86/contracts/token"
	"github.com/gin-gonic/gin"
)

type createContractRequest struct {
	Template db.Templates `json:"template" binding:"required,oneof=rental freelance services"`
}

type createContractResponse struct {
	Username string      `json:"username"`
	Contract db.Contract `json:"contract"`
	Party    string      `json:"party"`
}

func (server *Server) createContract(ctx *gin.Context) {
	var req createContractRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		server.Logger.Errorf("failed to unmarshal createContract request body")
		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.CreateContractParams{
		Template: req.Template,
		Username: authPayload.Username,
	}

	contract, err := server.store.CreateContract(ctx, arg)
	if err != nil {
		server.Logger.Errorf("failed to create contract")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	ownerURL := fmt.Sprintf("http://localhost:8080/contracts/%d/users/%s", contract.ID, authPayload.Username)

	rsp := createContractResponse{
		Username: authPayload.Username,
		Contract: contract,
		Party:    ownerURL,
	}

	ctx.JSON(http.StatusCreated, rsp)
}

type getContractRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getContract(ctx *gin.Context) {
	var req getContractRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	contract, err := server.store.GetContract(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err, http.StatusNotFound))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, contract)

}

type listContractRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listContract(ctx *gin.Context) {
	var req listContractRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.ListContractsParams{
		Username: authPayload.Username,
		Limit:    req.PageSize,
		Offset:   (req.PageID - 1) * req.PageSize,
	}

	contracts, err := server.store.ListContracts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, contracts)

}
