-- name: GetUserByLineID :one
SELECT * FROM users
WHERE line_id = $1;
