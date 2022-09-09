package db

import (
	"SimpleBank/utils"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	arg := CreateTransferParams{
		FromAccountID: sql.NullInt64{},
		ToAccountID:   sql.NullInt64{},
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {

}

func TestGetTransfer(t *testing.T) {

}

func TestUpdateTransfer(t *testing.T) {

}

func TestDeleteTransfer(t *testing.T) {

}

func TestListTransfer(t *testing.T) {

}
