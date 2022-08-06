// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
)

type Querier interface {
	CreateContract(ctx context.Context, arg CreateContractParams) (Contract, error)
	// CreateParty(ctx context.Context, arg CreatePartyParams) (Party, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	// CreateSignature(ctx context.Context, arg CreateSignatureParams) (Signature, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	// DeleteParty(ctx context.Context, arg DeletePartyParams) error
	// DeleteSignature(ctx context.Context, arg DeleteSignatureParams) error
	DeleteUser(ctx context.Context, username string) error
	GetContract(ctx context.Context, id string) (Contract, error)
	// GetParty(ctx context.Context, arg GetPartyParams) (Party, error)
	// GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	// GetSignature(ctx context.Context, arg GetSignatureParams) (Signature, error)
	GetUser(ctx context.Context, username string) (User, error)
	// ListContractParties(ctx context.Context, contractID int64) ([]Party, error)
	// ListContractSignatures(ctx context.Context, contractID int64) ([]Signature, error)
	// ListContracts(ctx context.Context, arg ListContractsParams) ([]Contract, error)
	// UpdateContract(ctx context.Context, arg UpdateContractParams) (Contract, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
