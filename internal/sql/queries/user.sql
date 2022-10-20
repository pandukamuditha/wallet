-- name: CreateUser :one
INSERT INTO "user" (fname, lname) VALUES ($1, $2) RETURNING id, fname, lname;

-- name: GetUser :one
SELECT id, fname, lname FROM "user" WHERE id = $1;

-- name: GetUserByFname :many
SELECT id, fname, lname FROM "user" WHERE fname = $1;