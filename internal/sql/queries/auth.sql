-- name: GetPasswordHash :one
SELECT password FROM "auth" WHERE username = $1;
 
-- name: GetUserIdByUsername :one
SELECT user_id FROM "auth" WHERE username = $1;
