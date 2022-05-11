package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomParty(t *testing.T) Party {
	user1 := createRandomUser(t)
	contract := createRandomContract(t)

	args := CreatePartyParams{
		UserID:     user1.ID,
		ContractID: contract.ID,
	}
	party, err := testQueries.CreateParty(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, party)

	require.Equal(t, user1.ID, party.UserID)
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
		UserID:     party1.UserID,
		ContractID: party1.ContractID,
	}

	party2, err := testQueries.GetParty(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, party2)

	require.Equal(t, party1.UserID, party2.UserID)
	require.Equal(t, party1.ContractID, party2.ContractID)
	require.WithinDuration(t, party1.CreatedAt, party2.CreatedAt, time.Second)

}

func TestDeleteParty(t *testing.T) {
	party1 := createRandomParty(t)

	arg := DeletePartyParams{
		UserID:     party1.UserID,
		ContractID: party1.ContractID,
	}

	err := testQueries.DeleteParty(context.Background(), arg)
	require.NoError(t, err)

	arg2 := GetPartyParams{
		UserID:     party1.UserID,
		ContractID: party1.ContractID,
	}

	party2, err := testQueries.GetParty(context.Background(), arg2)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, party2)
}

func TestListParties(t *testing.T) {
	var party Party
	for i := 0; i < 10; i++ {
		party = createRandomParty(t)
	}

	arg := ListPartiesParams{
		ContractID: party.ContractID,
		Limit:      1,
		Offset:     0,
	}

	parties, err := testQueries.ListParties(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, parties, 1)

	for _, contract := range parties {
		require.NotEmpty(t, contract)
	}
}
