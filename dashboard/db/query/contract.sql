-- name: CreateContract :one
INSERT INTO contracts (
  template
) VALUES (
  $1
)
RETURNING *;

-- name: GetContract :one
SELECT * FROM contracts
WHERE id = $1 LIMIT 1;

-- name: ListContracts :many
SELECT * FROM contracts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateContract :one
UPDATE contracts SET template = $2
WHERE id = $1
RETURNING *;

-- name: DeleteContract :exec
DELETE FROM contracts
WHERE id = $1;