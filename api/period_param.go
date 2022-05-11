package api

import (
	"database/sql"
	"net/http"

	db "github.com/danielmachado86/contracts/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createPeriodParamJSONRequest struct {
	Name  string `json:"name" binding:"required"`
	Value int32  `json:"value" binding:"required"`
	Units string `json:"units" binding:"required,oneof=days months years"`
}

type createPeriodParamUriRequest struct {
	ContractID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) createPeriodParam(ctx *gin.Context) {
	var JSONReq createPeriodParamJSONRequest
	if err := ctx.ShouldBindJSON(&JSONReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	var UriReq createPeriodParamUriRequest
	if err := ctx.ShouldBindUri(&UriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePeriodParamParams{
		ContractID: UriReq.ContractID,
		Name:       JSONReq.Name,
		Value:      JSONReq.Value,
		Units:      db.PeriodUnits(JSONReq.Units),
	}

	period, err := server.store.CreatePeriodParam(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, period)
}

type getPeriodParamRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPeriodParam(ctx *gin.Context) {
	var req getPeriodParamRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	contract, err := server.store.GetPeriodParam(ctx, req.ID)
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

type listPeriodParamRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type listPeriodParamUriRequest struct {
	ContractID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) listPeriodParam(ctx *gin.Context) {
	var req listPeriodParamRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var UriReq listPeriodParamUriRequest
	if err := ctx.ShouldBindQuery(&UriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPeriodParamsParams{
		Limit:      req.PageSize,
		Offset:     (req.PageID - 1) * req.PageSize,
		ContractID: UriReq.ContractID,
	}

	contracts, err := server.store.ListPeriodParams(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, contracts)

}
