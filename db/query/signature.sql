-- name: CreateSignature :one
INSERT INTO signatures (
  username,
  contract_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetSignature :one
SELECT * FROM signatures
WHERE username = $1 AND contract_id = $2 LIMIT 1;

-- name: ListContractSignatures :many
SELECT * FROM signatures
WHERE contract_id = $1
ORDER BY username
LIMIT NULL
OFFSET NULL;

-- name: DeleteSignature :exec
DELETE FROM signatures
WHERE username = $1 AND contract_id = $2;