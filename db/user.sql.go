// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  first_name, last_name, username, email, password_hashed
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING first_name, last_name, username, email, password_hashed, changed_at, created_at
`

type CreateUserParams struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	PasswordHashed string `json:"passwordHashed"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Username,
		arg.Email,
		arg.PasswordHashed,
	)
	var i User
	err := row.Scan(
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.PasswordHashed,
		&i.ChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, username)
	return err
}

const getUser = `-- name: GetUser :one
SELECT first_name, last_name, username, email, password_hashed, changed_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.PasswordHashed,
		&i.ChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET first_name = $2, last_name=$3
WHERE username = $1
RETURNING first_name, last_name, username, email, password_hashed, changed_at, created_at
`

type UpdateUserParams struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Username, arg.FirstName, arg.LastName)
	var i User
	err := row.Scan(
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.PasswordHashed,
		&i.ChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
