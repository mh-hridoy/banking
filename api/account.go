package api

import (
	db "github/mh-hridoy/banking/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAccountRequestParam struct {
	Owner    string `json:"owner" binding:"required" `
	Currency string `json:"currency" binding:"required,oneof=USD EUR" `
}

func (s *Server) CreateAccount(ctx *gin.Context) {
	var accountDetails CreateAccountRequestParam
	if err := ctx.ShouldBindJSON(&accountDetails); err != nil {
		ctx.JSON(http.StatusBadRequest, errorHandler(err))
		return
	}

	account, err := s.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    accountDetails.Owner,
		Currency: accountDetails.Currency,
		Balance:  0,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type GetSingleAccountParam struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (s *Server) GetSingleAccount(ctx *gin.Context) {
	var accountId GetSingleAccountParam

	err := ctx.ShouldBindUri(&accountId)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, errorHandler(err))
		return
	}

	account, errs := s.store.GetAccounts(ctx, int64(accountId.ID))

	if errs != nil {
		ctx.JSON(http.StatusBadRequest, errorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, account)

}

type ListAccountsParams struct {
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
	Offset int32 `form:"offset" binding:"required,min=1"`
}

func (s *Server) GetListOfAccount(ctx *gin.Context) {
	var listParams ListAccountsParams

	err := ctx.ShouldBindQuery(&listParams)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, errorHandler(err))
		return
	}

	accounts, errs := s.store.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  listParams.Limit,
		Offset: (listParams.Offset - 1) * listParams.Limit,
	})

	if errs != nil {
		ctx.JSON(http.StatusBadRequest, errorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)

}

type DeleteSingleAccountParam struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (s *Server) DeleteSingleAccount(ctx *gin.Context) {
	var accountId DeleteSingleAccountParam

	err := ctx.ShouldBindUri(&accountId)

	if err != nil {
		ctx.JSON(http.StatusNotAcceptable, errorHandler(err))
		return
	}

	errs := s.store.DeleteAccount(ctx, int64(accountId.ID))

	if errs != nil {
		ctx.JSON(http.StatusBadRequest, errorHandler(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Deleted Successfully"})

}
