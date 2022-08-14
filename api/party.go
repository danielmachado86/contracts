package api

// type createPartyRequest struct {
// 	ContractID int64 `uri:"id" binding:"required,min=1"`
// }
// type createPartyJSONRequest struct {
// 	Username string `json:"username" binding:"required"`
// }

// func (server *Server) createParty(ctx *gin.Context) {
// 	var req createPartyRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
// 		return
// 	}

// 	var JSONReq createPartyJSONRequest
// 	if err := ctx.ShouldBindJSON(&JSONReq); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
// 		return
// 	}

// 	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
// 	owner, err := server.store.GetContractOwner(ctx, req.ContractID)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err, http.StatusNotFound))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
// 		return
// 	}

// 	if owner.Username != authPayload.Username {
// 		err = errors.New("you are not the owner of the requested contract")
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err, http.StatusUnauthorized))
// 		return
// 	}

// 	arg := db.CreatePartyParams{
// 		Username:   JSONReq.Username,
// 		ContractID: req.ContractID,
// 		Role:       "signatory",
// 	}

// 	party, err := server.store.CreateParty(ctx, arg)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, party)
// }

// type getPartyRequest struct {
// 	Username   string `uri:"username" binding:"required"`
// 	ContractID int64  `uri:"id" binding:"required,min=1"`
// }

// func (server *Server) getParty(ctx *gin.Context) {
// 	var req getPartyRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
// 		return
// 	}

// 	arg := db.GetPartyParams{
// 		Username:   req.Username,
// 		ContractID: req.ContractID,
// 	}

// 	party, err := server.store.GetParty(ctx, arg)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err, http.StatusNotFound))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
// 		return
// 	}

// 	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
// 	if party.Username != authPayload.Username {
// 		err := errors.New("account doesn't belong to authenticated user")
// 		ctx.JSON(http.StatusUnauthorized, errorResponse(err, http.StatusUnauthorized))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, party)

// }

// type listPartiesRequest struct {
// 	PageID   int32 `form:"page_id" binding:"required,min=1"`
// 	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
// }

// type listPartiesURIRequest struct {
// 	ContractID int64 `uri:"id" binding:"required,min=1"`
// }

// func (server *Server) listParties(ctx *gin.Context) {
// 	var req listPartiesRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
// 		return
// 	}

// 	var UriReq listPartiesURIRequest
// 	if err := ctx.ShouldBindUri(&UriReq); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
// 		return
// 	}

// 	parties, err := server.store.ListContractParties(ctx, UriReq.ContractID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, parties)

// }

// type createSignatureRequest struct {
// 	ContractID int64 `uri:"id" binding:"required,min=1"`
// }

// func (server *Server) createSignature(ctx *gin.Context) {
// 	var req createSignatureRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
// 		return
// 	}

// 	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

// 	arg := db.GetPartyParams{
// 		Username:   authPayload.Username,
// 		ContractID: req.ContractID,
// 	}
// 	_, err := server.store.GetParty(ctx, arg)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			err = errors.New("you are not a party of the requested contract")
// 			ctx.JSON(http.StatusUnauthorized, errorResponse(err, http.StatusUnauthorized))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
// 		return
// 	}

// 	arg1 := db.CreateSignatureParams{
// 		Username:   authPayload.Username,
// 		ContractID: req.ContractID,
// 	}

// 	party, err := server.store.CreateSignature(ctx, arg1)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, party)
// }

// type listSignaturesURIRequest struct {
// 	ContractID int64 `uri:"id" binding:"required,min=1"`
// }

// func (server *Server) listSignatures(ctx *gin.Context) {

// 	var UriReq listSignaturesURIRequest
// 	if err := ctx.ShouldBindUri(&UriReq); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
// 		return
// 	}

// 	signatures, err := server.store.ListContractSignatures(ctx, UriReq.ContractID)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err, http.StatusInternalServerError))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, signatures)

// }
