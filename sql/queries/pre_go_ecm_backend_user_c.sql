-- name: GetUserByEmail :one
SELECT username, id FROM `users` WHERE username = ? LIMIT 1;

UPDATE `users`
SET 
    updated_at = $2
WHERE id = $1