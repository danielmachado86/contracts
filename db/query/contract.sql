-- name: CreateContract :one
WITH contracts AS (
  INSERT INTO contracts
    (template)
  VALUES
    ($1)
  RETURNING *),
parties AS (
  INSERT INTO parties
    (username, contract_id, role)
  SELECT $2, id, 'owner' FROM contracts
  RETURNING *
)
SELECT *
FROM contracts
LIMIT 1;

-- name: GetContract :one
SELECT * FROM contracts
WHERE id = $1 LIMIT 1;

-- name: ListContracts :many
SELECT contracts.* 
  FROM parties JOIN contracts ON parties.contract_id = contracts.id
  WHERE parties.username = $1
  ORDER BY contracts.created_at
  LIMIT $2
  OFFSET $3;

-- name: UpdateContract :one
UPDATE contracts SET template = $2
WHERE id = $1
RETURNING *;
