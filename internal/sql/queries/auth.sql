-- name: GetPasswordHash :one
SELECT password FROM "auth" WHERE username = $1;
 