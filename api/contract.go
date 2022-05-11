package api

import (
	"database/sql"
	"net/http"

	db "github.com/danielmachado86/contracts/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createContractRequest struct {
	Template string `json:"template" binding:"required,oneof=rental freelance services"`
}

func (server *Server) createContract(ctx *gin.Context) {
	var req createContractRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	contract, err := server.store.CreateContract(ctx, db.TemplatesRental)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, contract)
}

type getContractRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getContract(ctx *gin.Context) {
	var req getContractRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	contract, err := server.store.GetContract(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
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
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListContractsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	contracts, err := server.store.ListContracts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, contracts)

}