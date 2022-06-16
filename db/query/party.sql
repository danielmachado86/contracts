-- name: CreateParty :one
INSERT INTO parties (
  username,
  contract_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetParty :one
SELECT * FROM parties
WHERE username = $1 AND contract_id = $2 LIMIT 1;

-- name: ListParties :many
SELECT * FROM parties
WHERE contract_id = $1
ORDER BY username
LIMIT $2
OFFSET $3;

-- name: DeleteParty :exec
DELETE FROM parties
WHERE username = $1 AND contract_id = $2;