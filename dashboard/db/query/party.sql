-- name: CreateParty :one
INSERT INTO parties (
  user_id,
  contract_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetParty :one
SELECT * FROM parties
WHERE user_id = $1 AND contract_id = $2 LIMIT 1;

-- name: ListParties :many
SELECT * FROM parties
ORDER BY user_id, contract_id
LIMIT $1
OFFSET $2;

-- name: DeleteParty :exec
DELETE FROM parties
WHERE user_id = $1 AND contract_id = $2;