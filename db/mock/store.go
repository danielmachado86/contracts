// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/danielmachado86/contracts/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/danielmachado86/contracts/db"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateContract mocks base method.
func (m *MockStore) CreateContract(arg0 context.Context, arg1 db.CreateContractParams) (db.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContract", arg0, arg1)
	ret0, _ := ret[0].(db.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContract indicates an expected call of CreateContract.
func (mr *MockStoreMockRecorder) CreateContract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContract", reflect.TypeOf((*MockStore)(nil).CreateContract), arg0, arg1)
}

// CreateParty mocks base method.
func (m *MockStore) CreateParty(arg0 context.Context, arg1 db.CreatePartyParams) (db.Party, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateParty", arg0, arg1)
	ret0, _ := ret[0].(db.Party)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateParty indicates an expected call of CreateParty.
func (mr *MockStoreMockRecorder) CreateParty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateParty", reflect.TypeOf((*MockStore)(nil).CreateParty), arg0, arg1)
}

// CreatePeriodParam mocks base method.
func (m *MockStore) CreatePeriodParam(arg0 context.Context, arg1 db.CreatePeriodParamParams) (db.PeriodParam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePeriodParam", arg0, arg1)
	ret0, _ := ret[0].(db.PeriodParam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePeriodParam indicates an expected call of CreatePeriodParam.
func (mr *MockStoreMockRecorder) CreatePeriodParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePeriodParam", reflect.TypeOf((*MockStore)(nil).CreatePeriodParam), arg0, arg1)
}

// CreateSession mocks base method.
func (m *MockStore) CreateSession(arg0 context.Context, arg1 db.CreateSessionParams) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoreMockRecorder) CreateSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStore)(nil).CreateSession), arg0, arg1)
}

// CreateSignature mocks base method.
func (m *MockStore) CreateSignature(arg0 context.Context, arg1 db.CreateSignatureParams) (db.Signature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSignature", arg0, arg1)
	ret0, _ := ret[0].(db.Signature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSignature indicates an expected call of CreateSignature.
func (mr *MockStoreMockRecorder) CreateSignature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSignature", reflect.TypeOf((*MockStore)(nil).CreateSignature), arg0, arg1)
}

// CreateTimeParam mocks base method.
func (m *MockStore) CreateTimeParam(arg0 context.Context, arg1 db.CreateTimeParamParams) (db.TimeParam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTimeParam", arg0, arg1)
	ret0, _ := ret[0].(db.TimeParam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTimeParam indicates an expected call of CreateTimeParam.
func (mr *MockStoreMockRecorder) CreateTimeParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTimeParam", reflect.TypeOf((*MockStore)(nil).CreateTimeParam), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteContract mocks base method.
func (m *MockStore) DeleteContract(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContract", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContract indicates an expected call of DeleteContract.
func (mr *MockStoreMockRecorder) DeleteContract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContract", reflect.TypeOf((*MockStore)(nil).DeleteContract), arg0, arg1)
}

// DeleteParty mocks base method.
func (m *MockStore) DeleteParty(arg0 context.Context, arg1 db.DeletePartyParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteParty", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteParty indicates an expected call of DeleteParty.
func (mr *MockStoreMockRecorder) DeleteParty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteParty", reflect.TypeOf((*MockStore)(nil).DeleteParty), arg0, arg1)
}

// DeletePeriodParam mocks base method.
func (m *MockStore) DeletePeriodParam(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePeriodParam", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePeriodParam indicates an expected call of DeletePeriodParam.
func (mr *MockStoreMockRecorder) DeletePeriodParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePeriodParam", reflect.TypeOf((*MockStore)(nil).DeletePeriodParam), arg0, arg1)
}

// DeleteSignature mocks base method.
func (m *MockStore) DeleteSignature(arg0 context.Context, arg1 db.DeleteSignatureParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSignature", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSignature indicates an expected call of DeleteSignature.
func (mr *MockStoreMockRecorder) DeleteSignature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSignature", reflect.TypeOf((*MockStore)(nil).DeleteSignature), arg0, arg1)
}

// DeleteTimeParam mocks base method.
func (m *MockStore) DeleteTimeParam(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTimeParam", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTimeParam indicates an expected call of DeleteTimeParam.
func (mr *MockStoreMockRecorder) DeleteTimeParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTimeParam", reflect.TypeOf((*MockStore)(nil).DeleteTimeParam), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// GetContract mocks base method.
func (m *MockStore) GetContract(arg0 context.Context, arg1 int64) (db.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContract", arg0, arg1)
	ret0, _ := ret[0].(db.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContract indicates an expected call of GetContract.
func (mr *MockStoreMockRecorder) GetContract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContract", reflect.TypeOf((*MockStore)(nil).GetContract), arg0, arg1)
}

// GetContractOwner mocks base method.
func (m *MockStore) GetContractOwner(arg0 context.Context, arg1 int64) (db.Party, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContractOwner", arg0, arg1)
	ret0, _ := ret[0].(db.Party)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContractOwner indicates an expected call of GetContractOwner.
func (mr *MockStoreMockRecorder) GetContractOwner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContractOwner", reflect.TypeOf((*MockStore)(nil).GetContractOwner), arg0, arg1)
}

// GetParty mocks base method.
func (m *MockStore) GetParty(arg0 context.Context, arg1 db.GetPartyParams) (db.Party, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParty", arg0, arg1)
	ret0, _ := ret[0].(db.Party)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParty indicates an expected call of GetParty.
func (mr *MockStoreMockRecorder) GetParty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParty", reflect.TypeOf((*MockStore)(nil).GetParty), arg0, arg1)
}

// GetPeriodParam mocks base method.
func (m *MockStore) GetPeriodParam(arg0 context.Context, arg1 int64) (db.PeriodParam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPeriodParam", arg0, arg1)
	ret0, _ := ret[0].(db.PeriodParam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPeriodParam indicates an expected call of GetPeriodParam.
func (mr *MockStoreMockRecorder) GetPeriodParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeriodParam", reflect.TypeOf((*MockStore)(nil).GetPeriodParam), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStore) GetSession(arg0 context.Context, arg1 uuid.UUID) (db.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1)
	ret0, _ := ret[0].(db.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoreMockRecorder) GetSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStore)(nil).GetSession), arg0, arg1)
}

// GetSignature mocks base method.
func (m *MockStore) GetSignature(arg0 context.Context, arg1 db.GetSignatureParams) (db.Signature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSignature", arg0, arg1)
	ret0, _ := ret[0].(db.Signature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSignature indicates an expected call of GetSignature.
func (mr *MockStoreMockRecorder) GetSignature(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSignature", reflect.TypeOf((*MockStore)(nil).GetSignature), arg0, arg1)
}

// GetTimeParam mocks base method.
func (m *MockStore) GetTimeParam(arg0 context.Context, arg1 int64) (db.TimeParam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeParam", arg0, arg1)
	ret0, _ := ret[0].(db.TimeParam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTimeParam indicates an expected call of GetTimeParam.
func (mr *MockStoreMockRecorder) GetTimeParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeParam", reflect.TypeOf((*MockStore)(nil).GetTimeParam), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// ListContractParties mocks base method.
func (m *MockStore) ListContractParties(arg0 context.Context, arg1 int64) ([]db.Party, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContractParties", arg0, arg1)
	ret0, _ := ret[0].([]db.Party)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContractParties indicates an expected call of ListContractParties.
func (mr *MockStoreMockRecorder) ListContractParties(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContractParties", reflect.TypeOf((*MockStore)(nil).ListContractParties), arg0, arg1)
}

// ListContractSignatures mocks base method.
func (m *MockStore) ListContractSignatures(arg0 context.Context, arg1 int64) ([]db.Signature, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContractSignatures", arg0, arg1)
	ret0, _ := ret[0].([]db.Signature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContractSignatures indicates an expected call of ListContractSignatures.
func (mr *MockStoreMockRecorder) ListContractSignatures(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContractSignatures", reflect.TypeOf((*MockStore)(nil).ListContractSignatures), arg0, arg1)
}

// ListContracts mocks base method.
func (m *MockStore) ListContracts(arg0 context.Context, arg1 db.ListContractsParams) ([]db.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContracts", arg0, arg1)
	ret0, _ := ret[0].([]db.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContracts indicates an expected call of ListContracts.
func (mr *MockStoreMockRecorder) ListContracts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContracts", reflect.TypeOf((*MockStore)(nil).ListContracts), arg0, arg1)
}

// ListPeriodParams mocks base method.
func (m *MockStore) ListPeriodParams(arg0 context.Context, arg1 db.ListPeriodParamsParams) ([]db.PeriodParam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPeriodParams", arg0, arg1)
	ret0, _ := ret[0].([]db.PeriodParam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPeriodParams indicates an expected call of ListPeriodParams.
func (mr *MockStoreMockRecorder) ListPeriodParams(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPeriodParams", reflect.TypeOf((*MockStore)(nil).ListPeriodParams), arg0, arg1)
}

// ListTimeParams mocks base method.
func (m *MockStore) ListTimeParams(arg0 context.Context, arg1 db.ListTimeParamsParams) ([]db.TimeParam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTimeParams", arg0, arg1)
	ret0, _ := ret[0].([]db.TimeParam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTimeParams indicates an expected call of ListTimeParams.
func (mr *MockStoreMockRecorder) ListTimeParams(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTimeParams", reflect.TypeOf((*MockStore)(nil).ListTimeParams), arg0, arg1)
}

// UpdateContract mocks base method.
func (m *MockStore) UpdateContract(arg0 context.Context, arg1 db.UpdateContractParams) (db.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContract", arg0, arg1)
	ret0, _ := ret[0].(db.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateContract indicates an expected call of UpdateContract.
func (mr *MockStoreMockRecorder) UpdateContract(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContract", reflect.TypeOf((*MockStore)(nil).UpdateContract), arg0, arg1)
}

// UpdatePeriodParam mocks base method.
func (m *MockStore) UpdatePeriodParam(arg0 context.Context, arg1 db.UpdatePeriodParamParams) (db.PeriodParam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePeriodParam", arg0, arg1)
	ret0, _ := ret[0].(db.PeriodParam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePeriodParam indicates an expected call of UpdatePeriodParam.
func (mr *MockStoreMockRecorder) UpdatePeriodParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePeriodParam", reflect.TypeOf((*MockStore)(nil).UpdatePeriodParam), arg0, arg1)
}

// UpdateTimeParam mocks base method.
func (m *MockStore) UpdateTimeParam(arg0 context.Context, arg1 db.UpdateTimeParamParams) (db.TimeParam, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTimeParam", arg0, arg1)
	ret0, _ := ret[0].(db.TimeParam)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTimeParam indicates an expected call of UpdateTimeParam.
func (mr *MockStoreMockRecorder) UpdateTimeParam(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTimeParam", reflect.TypeOf((*MockStore)(nil).UpdateTimeParam), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
