// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
)

type Querier interface {
	CreateContract(ctx context.Context, arg CreateContractParams) (Contract, error)
	CreateParty(ctx context.Context, arg CreatePartyParams) (Party, error)
	CreatePeriodParam(ctx context.Context, arg CreatePeriodParamParams) (PeriodParam, error)
	CreateTimeParam(ctx context.Context, arg CreateTimeParamParams) (TimeParam, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteContract(ctx context.Context, id int64) error
	DeleteParty(ctx context.Context, arg DeletePartyParams) error
	DeletePeriodParam(ctx context.Context, id int64) error
	DeleteTimeParam(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, username string) error
	GetContract(ctx context.Context, id int64) (Contract, error)
	GetContractOwner(ctx context.Context, contractID int64) (Party, error)
	GetParty(ctx context.Context, arg GetPartyParams) (Party, error)
	GetPeriodParam(ctx context.Context, id int64) (PeriodParam, error)
	GetTimeParam(ctx context.Context, id int64) (TimeParam, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListContractParties(ctx context.Context, arg ListContractPartiesParams) ([]Party, error)
	ListContracts(ctx context.Context, arg ListContractsParams) ([]Contract, error)
	ListPeriodParams(ctx context.Context, arg ListPeriodParamsParams) ([]PeriodParam, error)
	ListTimeParams(ctx context.Context, arg ListTimeParamsParams) ([]TimeParam, error)
	UpdateContract(ctx context.Context, arg UpdateContractParams) (Contract, error)
	UpdatePeriodParam(ctx context.Context, arg UpdatePeriodParamParams) (PeriodParam, error)
	UpdateTimeParam(ctx context.Context, arg UpdateTimeParamParams) (TimeParam, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
