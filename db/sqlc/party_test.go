package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/danielmachado86/contracts/utils"
	"github.com/stretchr/testify/require"
)

func createRandomParty(t *testing.T) Party {
	user1 := createRandomUser(t)
	contract := createRandomContract(t)
	role := utils.RandomRole()

	args := CreatePartyParams{
		Username:   user1.Username,
		ContractID: contract.ID,
		Role:       ContractRole(role),
	}
	party, err := testQueries.CreateParty(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, party)

	require.Equal(t, user1.Username, party.Username)
	require.Equal(t, contract.ID, party.ContractID)
	require.NotZero(t, party.CreatedAt)

	return party
}

func TestCreateParty(t *testing.T) {
	createRandomParty(t)
}

func TestGetParty(t *testing.T) {
	party1 := createRandomParty(t)

	args := GetPartyParams{
		Username:   party1.Username,
		ContractID: party1.ContractID,
	}

	party2, err := testQueries.GetParty(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, party2)

	require.Equal(t, party1.Username, party2.Username)
	require.Equal(t, party1.ContractID, party2.ContractID)
	require.WithinDuration(t, party1.CreatedAt, party2.CreatedAt, time.Second)

}

func TestDeleteParty(t *testing.T) {
	party1 := createRandomParty(t)

	arg := DeletePartyParams{
		Username:   party1.Username,
		ContractID: party1.ContractID,
	}

	err := testQueries.DeleteParty(context.Background(), arg)
	require.NoError(t, err)

	arg2 := GetPartyParams{
		Username:   party1.Username,
		ContractID: party1.ContractID,
	}

	party2, err := testQueries.GetParty(context.Background(), arg2)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, party2)
}
