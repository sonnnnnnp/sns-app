-- name: GetUserByName :one
SELECT * FROM users
WHERE name = @name::text;
