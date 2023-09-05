package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {

	store := NewStore(testDb)

	account1 := createRandomAccount()
	account2 := createRandomAccount()

	fmt.Println("balance is", account1.Balance, account2.Balance)

	n := 5

	amount := int64(10)

	errors := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {

		go func() {
			result, err := store.TransferExec(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errors <- err
			results <- result
		}()

	}

	for i := 0; i < n; i++ {

		err := <-errors
		require.NoError(t, err)

		result := <-results

		require.NotEmpty(t, result)

		transfer := result.Transfer
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)

		require.Equal(t, transfer.Amount, amount)

		fromAccount, err := store.GetAccounts(context.Background(), account1.ID)

		require.NoError(t, err)

		require.NotEmpty(t, fromAccount)
		require.Equal(t, fromAccount.ID, account1.ID)
		require.Equal(t, fromAccount.ID, account1.ID)

		toAccount, err := store.GetAccounts(context.Background(), account2.ID)

		require.NoError(t, err)

		require.NotEmpty(t, toAccount)
		require.Equal(t, toAccount.ID, account2.ID)
		require.Equal(t, toAccount.ID, account2.ID)

		fromEntry, err := store.GetEntry(context.Background(), result.FromEntry.ID)
		require.NoError(t, err)
		require.NotEmpty(t, fromEntry)
		require.Equal(t, fromEntry.Amount, result.FromEntry.Amount)

		toEntry, err := store.GetEntry(context.Background(), result.ToEntry.ID)
		require.NoError(t, err)
		require.NotEmpty(t, toEntry)
		require.Equal(t, toEntry.Amount, result.ToEntry.Amount)

		transferDetails, err := store.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err)
		require.NotEmpty(t, transferDetails)
		require.Equal(t, transferDetails.Amount, result.Transfer.Amount)

		fmt.Println(">>tx", fromAccount.Balance, toAccount.Balance)

	}

}
