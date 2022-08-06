// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: party.sql

package db

import (
	"context"
)

const createParty = `-- name: CreateParty :one
INSERT INTO parties (
  username,
  contract_id,
  role
) VALUES (
  $1, $2, $3
)
RETURNING username, role, contract_id, created_at
`

type CreatePartyParams struct {
	Username   string `json:"username"`
	ContractID string  `json:"contractID"`
	Role       string `json:"role"`
}

func (q *Queries) CreateParty(ctx context.Context, arg CreatePartyParams) (Party, error) {
	row := q.db.QueryRowContext(ctx, createParty, arg.Username, arg.ContractID, arg.Role)
	var i Party
	err := row.Scan(
		&i.Username,
		&i.Role,
		&i.ContractID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteParty = `-- name: DeleteParty :exec
DELETE FROM parties
WHERE username = $1 AND contract_id = $2
`

type DeletePartyParams struct {
	Username   string `json:"username"`
	ContractID int64  `json:"contractID"`
}

func (q *Queries) DeleteParty(ctx context.Context, arg DeletePartyParams) error {
	_, err := q.db.ExecContext(ctx, deleteParty, arg.Username, arg.ContractID)
	return err
}

const getParty = `-- name: GetParty :one
SELECT username, role, contract_id, created_at FROM parties
WHERE username = $1 AND contract_id = $2 LIMIT 1
`

type GetPartyParams struct {
	Username   string `json:"username"`
	ContractID int64  `json:"contractID"`
}

func (q *Queries) GetParty(ctx context.Context, arg GetPartyParams) (Party, error) {
	row := q.db.QueryRowContext(ctx, getParty, arg.Username, arg.ContractID)
	var i Party
	err := row.Scan(
		&i.Username,
		&i.Role,
		&i.ContractID,
		&i.CreatedAt,
	)
	return i, err
}

const listContractParties = `-- name: ListContractParties :many
SELECT username, role, contract_id, created_at FROM parties
WHERE contract_id = $1
ORDER BY username
LIMIT NULL
OFFSET NULL
`

func (q *Queries) ListContractParties(ctx context.Context, contractID int64) ([]Party, error) {
	rows, err := q.db.QueryContext(ctx, listContractParties, contractID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Party{}
	for rows.Next() {
		var i Party
		if err := rows.Scan(
			&i.Username,
			&i.Role,
			&i.ContractID,
			&i.CreatedAt,
		); err != nil {
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
