-- name: CreateParty :one
INSERT INTO parties (
  username,
  contract_id,
  role
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetParty :one
SELECT * FROM parties
WHERE username = $1 AND contract_id = $2 LIMIT 1;

-- name: GetContractOwner :one
SELECT * FROM parties
WHERE contract_id = $1 AND role = 'owner' LIMIT 1;

-- name: ListContractParties :many
SELECT * FROM parties
WHERE contract_id = $1
ORDER BY username
LIMIT NULL
OFFSET NULL;

-- name: DeleteParty :exec
DELETE FROM parties
WHERE username = $1 AND contract_id = $2;