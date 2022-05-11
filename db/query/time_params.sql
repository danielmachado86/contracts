-- name: CreateTimeParam :one
INSERT INTO time_params (
  contract_id,
  name,
  value
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTimeParam :one
SELECT * FROM time_params
WHERE id = $1 LIMIT 1;

-- name: ListTimeParams :many
SELECT * FROM time_params
WHERE contract_id = $1
ORDER BY value
LIMIT $2
OFFSET $3;

-- name: UpdateTimeParam :one
UPDATE time_params SET value = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTimeParam :exec
DELETE FROM time_params
WHERE id = $1;