package api

import (
	"database/sql"
	"net/http"

	db "github.com/danielmachado86/contracts/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createPartyRequest struct {
	UserID     int64 `uri:"userID" binding:"required,min=1"`
	ContractID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) createParty(ctx *gin.Context) {
	var req createPartyRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreatePartyParams{
		UserID:     req.UserID,
		ContractID: req.ContractID,
	}

	party, err := server.store.CreateParty(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, party)
}

type getPartyRequest struct {
	UserID     int64 `uri:"userID" binding:"required,min=1"`
	ContractID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getParty(ctx *gin.Context) {
	var req getPartyRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetPartyParams{
		UserID:     req.UserID,
		ContractID: req.ContractID,
	}

	party, err := server.store.GetParty(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, party)

}

type listPartiesRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type listPartiesURIRequest struct {
	ContractID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) listParties(ctx *gin.Context) {
	var req listPartiesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var UriReq listPartiesURIRequest
	if err := ctx.ShouldBindUri(&UriReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListPartiesParams{
		ContractID: UriReq.ContractID,
		Limit:      req.PageSize,
		Offset:     (req.PageID - 1) * req.PageSize,
	}

	parties, err := server.store.ListParties(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, parties)

}
