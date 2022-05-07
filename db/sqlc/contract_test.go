package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomContract(t *testing.T) Contract {
	contract, err := testQueries.CreateContract(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, contract)

	require.Equal(t, int32(1), contract.Template)

	require.NotZero(t, contract.ID)

	return contract
}

func TestCreateContract(t *testing.T) {
	createRandomContract(t)
}

func TestGetContract(t *testing.T) {
	contract1 := createRandomContract(t)
	contract2, err := testQueries.GetContract(context.Background(), contract1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, contract2)

	require.Equal(t, contract1.ID, contract2.ID)

}

func TestUpdateContract(t *testing.T) {
	contract1 := createRandomContract(t)

	arg := UpdateContractParams{
		ID:       contract1.ID,
		Template: 3,
	}

	contract2, err := testQueries.UpdateContract(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, contract2)

	require.Equal(t, contract1.ID, contract2.ID)
	require.Equal(t, arg.Template, contract2.Template)

}

func TestDeleteContract(t *testing.T) {
	contract1 := createRandomContract(t)

	err := testQueries.DeleteContract(context.Background(), contract1.ID)
	require.NoError(t, err)

	contract2, err := testQueries.GetContract(context.Background(), contract1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, contract2)
}

func TestListContract(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomContract(t)
	}

	arg := ListContractsParams{
		Limit:  5,
		Offset: 5,
	}

	contracts, err := testQueries.ListContracts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, contracts, 5)

	for _, contract := range contracts {
		require.NotEmpty(t, contract)
	}
}
