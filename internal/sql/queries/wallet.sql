-- name: CreateWallet :one
INSERT INTO "wallet" (user_id) VALUES (?) RETURNING id, user_id;

-- name: CreateWalletWithBalance :one
INSERT INTO "wallet" (user_id, balance) VALUES (?, ?) RETURNING id, user_id, balance;

-- name: GetWalletById :one
SELECT id, user_id, balance FROM "wallet" WHERE id = ?;

-- name: GetWalletByUserId :one
SELECT id, user_id, balance FROM "wallet" WHERE user_id = ?;