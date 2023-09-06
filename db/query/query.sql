-- name: GetAccounts :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountsForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id LIMIT $1 OFFSET $2;

-- name: CreateAccount :one
INSERT INTO accounts (
  owner, balance, currency, created_at
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;

-- name: UpdateAccount :exec
UPDATE accounts
  set balance = $2
WHERE id = $1;


-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id;


-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;

-- name: UpdateTransfer :exec
UPDATE transfers
  set amount = $2
WHERE id = $1;


-- name: CreateEntry :one
INSERT INTO entries (
  account_id, amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
ORDER BY id;


-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1;

-- name: UpdateEntry :exec
UPDATE entries
  set amount = $2
WHERE id = $1;
