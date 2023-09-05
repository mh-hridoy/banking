package db

import (
	"context"
	"github/mh-hridoy/banking/utils"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	arg := CreateAccountParams{
		Owner:     utils.GetOwnerName(),
		Balance:   utils.GenerateBalance(),
		Currency:  utils.GenerateCurrency(),
		CreatedAt: time.Now(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner, arg.Balance, account.Balance, arg.Currency, account.Currency)
	require.Equal(t, arg.CreatedAt.Local(), account.CreatedAt.Local())

}

func createRandomAccount() *Account {
	result, err := testQueries.CreateAccount(context.Background(), CreateAccountParams{
		Owner:    utils.GetOwnerName(),
		Balance:  utils.GenerateBalance(),
		Currency: utils.GenerateCurrency(),
	})

	if err != nil {
		log.Fatal(err)
	}
	return &Account{
		Owner:    result.Owner,
		Balance:  result.Balance,
		Currency: result.Currency,
		ID:       result.ID,
	}
}
