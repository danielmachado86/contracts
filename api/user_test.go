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
	"time"

	mockdb "github.com/danielmachado86/contracts/db/mock"
	db "github.com/danielmachado86/contracts/db/sqlc"
	"github.com/danielmachado86/contracts/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer()
	require.NoError(t, err)

	err = server.ConfigServer(config, store)
	require.NoError(t, err)

	return server
}

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}
	err := utils.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}
	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestCreateUserRequiredParams(t *testing.T) {
	user, password := randomUser(t)

	tt := []struct {
		name          string
		body          gin.H
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "EmptyBody",
			body: gin.H{},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				//TODO: Check error message
			},
		},
		{
			name: "NameMissing",
			body: gin.H{
				"lastName": user.LastName,
				"username": user.Username,
				"email":    user.Email,
				"password": password,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				//TODO: Check error message
			},
		},
		{
			name: "NameIsBlank",
			body: gin.H{
				"Name":     "",
				"lastName": user.LastName,
				"username": user.Username,
				"email":    user.Email,
				"password": password,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				//TODO: Check error message
			},
		},
		{
			name: "LastNameMissing",
			body: gin.H{
				"Name":     user.Name,
				"username": user.Username,
				"email":    user.Email,
				"password": password,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				//TODO: Check error message
			},
		},
		{
			name: "LastNameIsBlank",
			body: gin.H{
				"Name":     user.Name,
				"LastName": "",
				"username": user.Username,
				"email":    user.Email,
				"password": password,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				//TODO: Check error message
			},
		},
		{
			name: "UsernameMissing",
			body: gin.H{
				"Name":     user.Name,
				"LastName": user.LastName,
				"email":    user.Email,
				"password": password,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				//TODO: Check error message
			},
		},
		{
			name: "UsernameIsBlank",
			body: gin.H{
				"Name":     user.Name,
				"LastName": user.LastName,
				"username": "",
				"email":    user.Email,
				"password": password,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				//TODO: Check error message
			},
		},
		{
			name: "EmailMissing",
			body: gin.H{
				"Name":     user.Name,
				"LastName": user.LastName,
				"Username": user.Username,
				"password": password,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				//TODO: Check error message
			},
		},
		{
			name: "EmailIsBlank",
			body: gin.H{
				"Name":     user.Name,
				"LastName": user.LastName,
				"Username": user.Username,
				"email":    "",
				"password": password,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				//TODO: Check error message
			},
		},
		{
			name: "PaswordHashMissing",
			body: gin.H{
				"Name":     user.Name,
				"LastName": user.LastName,
				"Username": user.Username,
				"email":    user.Email,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				//TODO: Check error message
			},
		},
		{
			name: "ValidEmail",
			body: gin.H{
				"Name":     user.Name,
				"LastName": user.LastName,
				"Username": user.Username,
				"email":    "invalid email",
				"password": password,
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
				//TODO: Check error message
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			store := mockdb.NewMockStore(ctrl)

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

func TestUserMockDB(t *testing.T) {
	user, password := randomUser(t)
	hashedPassword, err := utils.HashPasword(password)
	require.NoError(t, err)

	tt := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":     user.Name,
				"lastName": user.LastName,
				"username": user.Username,
				"email":    user.Email,
				"password": password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateUserParams{
					Name:           user.Name,
					LastName:       user.LastName,
					Username:       user.Username,
					Email:          user.Email,
					HashedPassword: hashedPassword,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"name":     user.Name,
				"lastName": user.LastName,
				"username": user.Username,
				"email":    user.Email,
				"password": password,
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
				"lastName": user.LastName,
				"username": user.Username,
				"email":    user.Email,
				"password": password,
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

func TestUserEnpoints(t *testing.T) {
	tt := []struct {
		name          string
		url           string
		method        string
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "UsersPOSTShouldExists",
			url:    "/users",
			method: "POST",
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.NotEqual(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "UsersDELETEShouldNotExists",
			url:    "/users",
			method: "DELETE",
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "UsersPUTShouldNotExists",
			url:    "/users",
			method: "PUT",
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			store := mockdb.NewMockStore(ctrl)

			server := newTestServer(t, store)
			rec := httptest.NewRecorder()

			req, err := http.NewRequest(tc.method, tc.url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(rec, req)
			tc.checkResponse(rec)

		})
	}
}

func randomUser(t *testing.T) (user db.User, password string) {
	password = utils.RandomString(6)
	hashedPassword, err := utils.HashPasword(password)
	require.NoError(t, err)

	user = db.User{
		Name:           utils.RandomString(6),
		LastName:       utils.RandomString(6),
		Username:       utils.RandomString(6),
		Email:          utils.RandomEmail(),
		HashedPassword: hashedPassword,
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
	require.Empty(t, gotUser.HashedPassword)
}
