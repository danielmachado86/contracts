// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: contract.sql

package db

import (
	"context"
)

const createContract = `-- name: CreateContract :one
WITH contracts AS (
  INSERT INTO contracts
    (template)
  VALUES
    ($1)
  RETURNING id, template, created_at),
parties AS (
  INSERT INTO parties
    (username, contract_id, role)
  SELECT $2, id, 'owner' FROM contracts
  RETURNING username, role, contract_id, created_at
)
SELECT id, template, created_at
FROM contracts
LIMIT 1
`

type CreateContractParams struct {
	Template Templates `json:"template"`
	Username string    `json:"username"`
}

func (q *Queries) CreateContract(ctx context.Context, arg CreateContractParams) (Contract, error) {
	row := q.db.QueryRowContext(ctx, createContract, arg.Template, arg.Username)
	var i Contract
	err := row.Scan(&i.ID, &i.Template, &i.CreatedAt)
	return i, err
}

const deleteContract = `-- name: DeleteContract :exec
DELETE FROM contracts
WHERE id = $1
`

func (q *Queries) DeleteContract(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteContract, id)
	return err
}

const getContract = `-- name: GetContract :one
SELECT id, template, created_at FROM contracts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetContract(ctx context.Context, id int64) (Contract, error) {
	row := q.db.QueryRowContext(ctx, getContract, id)
	var i Contract
	err := row.Scan(&i.ID, &i.Template, &i.CreatedAt)
	return i, err
}

const listContracts = `-- name: ListContracts :many
SELECT contracts.id, contracts.template, contracts.created_at 
  FROM parties JOIN contracts ON parties.contract_id = contracts.id
  WHERE parties.username = $1
  ORDER BY contracts.created_at
  LIMIT $2
  OFFSET $3
`

type ListContractsParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *Queries) ListContracts(ctx context.Context, arg ListContractsParams) ([]Contract, error) {
	rows, err := q.db.QueryContext(ctx, listContracts, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Contract{}
	for rows.Next() {
		var i Contract
		if err := rows.Scan(&i.ID, &i.Template, &i.CreatedAt); err != nil {
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

const updateContract = `-- name: UpdateContract :one
UPDATE contracts SET template = $2
WHERE id = $1
RETURNING id, template, created_at
`

type UpdateContractParams struct {
	ID       int64     `json:"id"`
	Template Templates `json:"template"`
}

func (q *Queries) UpdateContract(ctx context.Context, arg UpdateContractParams) (Contract, error) {
	row := q.db.QueryRowContext(ctx, updateContract, arg.ID, arg.Template)
	var i Contract
	err := row.Scan(&i.ID, &i.Template, &i.CreatedAt)
	return i, err
}
