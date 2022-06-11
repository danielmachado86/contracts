package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/danielmachado86/contracts/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Name:         utils.RandomString(6),
		LastName:     utils.RandomString(6),
		Username:     utils.RandomUser(),
		Email:        utils.RandomEmail(),
		PasswordHash: "password",
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)

	require.NotZero(t, user.ID)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)

}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		ID:   user1.ID,
		Name: "Daniel",
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.NotZero(t, user2.CreatedAt)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Name, user2.Name)

}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user1.ID)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, contract := range users {
		require.NotEmpty(t, contract)
	}
}
