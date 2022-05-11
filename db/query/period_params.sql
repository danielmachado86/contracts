-- name: CreatePeriodParam :one
INSERT INTO period_params (
  contract_id,
  name,
  value,
  units
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetPeriodParam :one
SELECT * FROM period_params
WHERE id = $1 LIMIT 1;

-- name: ListPeriodParams :many
SELECT * FROM period_params
WHERE contract_id = $1
ORDER BY value
LIMIT $2
OFFSET $3;

-- name: UpdatePeriodParam :one
UPDATE period_params SET value = $2
WHERE id = $1
RETURNING *;

-- name: DeletePeriodParam :exec
DELETE FROM period_params
WHERE id = $1;