// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  name, last_name, username, email, hashed_password
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING name, last_name, username, email, hashed_password, password_changed_at, created_at
`

type CreateUserParams struct {
	Name           string `json:"name"`
	LastName       string `json:"lastName"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashedPassword"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Name,
		arg.LastName,
		arg.Username,
		arg.Email,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.Name,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
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
SELECT name, last_name, username, email, hashed_password, password_changed_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Name,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users SET name = $2
WHERE username = $1
RETURNING name, last_name, username, email, hashed_password, password_changed_at, created_at
`

type UpdateUserParams struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Username, arg.Name)
	var i User
	err := row.Scan(
		&i.Name,
		&i.LastName,
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
