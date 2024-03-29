// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: auth.sql

package repositories

import (
	"context"
)

const getPasswordHash = `-- name: GetPasswordHash :one
SELECT password FROM "auth" WHERE username = $1
`

func (q *Queries) GetPasswordHash(ctx context.Context, username string) (string, error) {
	row := q.db.QueryRow(ctx, getPasswordHash, username)
	var password string
	err := row.Scan(&password)
	return password, err
}

const getUserIdByUsername = `-- name: GetUserIdByUsername :one
SELECT user_id FROM "auth" WHERE username = $1
`

func (q *Queries) GetUserIdByUsername(ctx context.Context, username string) (int64, error) {
	row := q.db.QueryRow(ctx, getUserIdByUsername, username)
	var user_id int64
	err := row.Scan(&user_id)
	return user_id, err
}
