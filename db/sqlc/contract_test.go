package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/danielmachado86/contracts/utils"
	"github.com/stretchr/testify/require"
)

func createRandomContract(t *testing.T) Contract {
	user := createRandomUser(t)
	arg := CreateContractParams{
		Template: TemplatesRental,
		Username: user.Username,
	}
	contract, err := testQueries.CreateContract(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, contract)

	require.Equal(t, TemplatesRental, contract.Template)

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
		Template: TemplatesRental,
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

	arg0 := CreateUserParams{
		Name:           utils.RandomString(6),
		LastName:       utils.RandomString(6),
		Username:       "username",
		Email:          utils.RandomEmail(),
		HashedPassword: "password",
	}
	user, err := testQueries.CreateUser(context.Background(), arg0)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	arg1 := ListContractsParams{
		Username: "username",
		Limit:    5,
		Offset:   5,
	}

	_, err = testFailingQueries.ListContracts(context.Background(), arg1)
	require.Error(t, err)

	for i := 0; i < 10; i++ {
		arg := CreateContractParams{
			Template: Templates(utils.RandomTemplate()),
			Username: "username",
		}

		contract, err := testQueries.CreateContract(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, contract)
	}

	contracts, err := testQueries.ListContracts(context.Background(), arg1)
	require.NoError(t, err)
	require.Len(t, contracts, 5)

	for _, contract := range contracts {
		require.NotEmpty(t, contract)
	}
}
