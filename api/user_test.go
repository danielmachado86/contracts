package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	mockdb "github.com/danielmachado86/contracts/db/mock"
	db "github.com/danielmachado86/contracts/db/sqlc"
	"github.com/danielmachado86/contracts/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	server, err := NewServer(store)
	require.NoError(t, err)

	return server
}

type eqCreateUserParamsMatcher struct {
	arg db.CreateUserParams
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v", e.arg)
}

func EqCreateUserParams(arg db.CreateUserParams) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg}
}

func TestCreateUser(t *testing.T) {
	user := randomUser(t)

	tt := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":         user.Name,
				"lastName":     user.LastName,
				"username":     user.Username,
				"email":        user.Email,
				"passwordHash": user.PasswordHash,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateUserParams{
					Name:         user.Name,
					LastName:     user.LastName,
					Username:     user.Username,
					Email:        user.Email,
					PasswordHash: user.PasswordHash,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), arg).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"name":         user.Name,
				"lastName":     user.LastName,
				"username":     user.Username,
				"email":        user.Email,
				"passwordHash": user.PasswordHash,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "MissingField",
			body: gin.H{
				"lastName":     user.LastName,
				"username":     user.Username,
				"email":        user.Email,
				"passwordHash": user.PasswordHash,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			rec := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			rPath := "/users"
			req, err := http.NewRequest(http.MethodPost, rPath, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(rec, req)
			tc.checkResponse(rec)

		})
	}
}

func randomUser(t *testing.T) (user db.User) {
	user = db.User{
		Name:         utils.RandomString(6),
		LastName:     utils.RandomString(6),
		Username:     utils.RandomString(6),
		Email:        utils.RandomEmail(),
		PasswordHash: "password",
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.Name, gotUser.Name)
	require.Equal(t, user.LastName, gotUser.LastName)
	require.Equal(t, user.Email, gotUser.Email)
	require.Equal(t, user.PasswordHash, gotUser.PasswordHash)
}
