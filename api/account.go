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
