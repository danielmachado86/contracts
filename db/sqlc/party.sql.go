// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: party.sql

package db

import (
	"context"
)

const createParty = `-- name: CreateParty :one
INSERT INTO parties (
  user_id,
  contract_id
) VALUES (
  $1, $2
)
RETURNING user_id, contract_id, created_at
`

type CreatePartyParams struct {
	UserID     int64 `json:"userID"`
	ContractID int64 `json:"contractID"`
}

func (q *Queries) CreateParty(ctx context.Context, arg CreatePartyParams) (Party, error) {
	row := q.db.QueryRowContext(ctx, createParty, arg.UserID, arg.ContractID)
	var i Party
	err := row.Scan(&i.UserID, &i.ContractID, &i.CreatedAt)
	return i, err
}

const deleteParty = `-- name: DeleteParty :exec
DELETE FROM parties
WHERE user_id = $1 AND contract_id = $2
`

type DeletePartyParams struct {
	UserID     int64 `json:"userID"`
	ContractID int64 `json:"contractID"`
}

func (q *Queries) DeleteParty(ctx context.Context, arg DeletePartyParams) error {
	_, err := q.db.ExecContext(ctx, deleteParty, arg.UserID, arg.ContractID)
	return err
}

const getParty = `-- name: GetParty :one
SELECT user_id, contract_id, created_at FROM parties
WHERE user_id = $1 AND contract_id = $2 LIMIT 1
`

type GetPartyParams struct {
	UserID     int64 `json:"userID"`
	ContractID int64 `json:"contractID"`
}

func (q *Queries) GetParty(ctx context.Context, arg GetPartyParams) (Party, error) {
	row := q.db.QueryRowContext(ctx, getParty, arg.UserID, arg.ContractID)
	var i Party
	err := row.Scan(&i.UserID, &i.ContractID, &i.CreatedAt)
	return i, err
}

const listParties = `-- name: ListParties :many
SELECT user_id, contract_id, created_at FROM parties
ORDER BY user_id, contract_id
LIMIT $1
OFFSET $2
`

type ListPartiesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListParties(ctx context.Context, arg ListPartiesParams) ([]Party, error) {
	rows, err := q.db.QueryContext(ctx, listParties, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Party{}
	for rows.Next() {
		var i Party
		if err := rows.Scan(&i.UserID, &i.ContractID, &i.CreatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
