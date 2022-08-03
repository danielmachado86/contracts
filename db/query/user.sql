-- name: CreateUser :one
INSERT INTO users (
  first_name, last_name, username, email, password_hashed
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users SET first_name = $2, last_name=$3
WHERE username = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;