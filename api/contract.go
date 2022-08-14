package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/danielmachado86/contracts/db"
	"github.com/danielmachado86/contracts/rules"
	"github.com/danielmachado86/contracts/token"
	"github.com/gin-gonic/gin"
)

type createContractRequest struct {
	Name        string `json:"name" dynamodbav:"name" binding:"required"`
	Template    string `json:"template" binding:"required,oneof=rental freelance services"`
	Description string `json:"description" dynamodbav:"description"`
}

type createContractResponse struct {
	Username    string      `json:"username"`
	Contract    db.Contract `json:"contract"`
	ContractUrl string      `json:"contract_url"`
	Party       string      `json:"party"`
}

func (server *Server) createContract(ctx *gin.Context) {
	var req createContractRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		server.Logger.Errorf("failed to unmarshal createContract request body")
		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	user, err := server.store.GetUser(ctx, authPayload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err, http.StatusNotFound))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	party := db.PartyView{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	arg := db.CreateContractParams{
		Name:        req.Name,
		Description: req.Description,
		Template:    req.Template,
		Owner:       party,
	}

	contract, err := server.store.CreateContract(ctx, arg)
	if err != nil {
		server.Logger.Errorf("failed to create contract")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	var meta map[string]interface{}
	mContract, err := json.Marshal(contract)
	if err != nil {
		server.Logger.Errorf("failed to marshal contract")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}
	err = json.Unmarshal(mContract, &meta)
	if err != nil {
		server.Logger.Errorf("failed to unmarshal contract")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	err = rules.CalculateTerms(ctx, server.store, meta)
	if err != nil {
		server.Logger.Errorf("failed to calculate contract")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
		return
	}

	contractUrl := fmt.Sprintf("http://localhost:8080/contracts/%s", contract.ID)
	ownerURL := fmt.Sprintf("http://localhost:8080/contracts/%s/users/%s", contract.ID, authPayload.Username)

	rsp := createContractResponse{
		Username:    authPayload.Username,
		Contract:    contract,
		ContractUrl: contractUrl,
		Party:       ownerURL,
	}

	ctx.JSON(http.StatusCreated, rsp)
}

type getContractURIRequest struct {
	ID string `uri:"id"`
}

func (server *Server) getContract(ctx *gin.Context) {
	var req getContractURIRequest
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

// type listContractRequest struct {
// 	PageID   int32 `form:"page_id" binding:"required,min=1"`
// 	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
// }

// func (server *Server) listContract(ctx *gin.Context) {
// 	var req listContractRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
// 		return
// 	}

// 	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

// 	arg := db.ListContractsParams{
// 		Username: authPayload.Username,
// 		Limit:    req.PageSize,
// 		Offset:   (req.PageID - 1) * req.PageSize,
// 	}

// 	contracts, err := server.store.ListContracts(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, contracts)

// }
