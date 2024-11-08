-- name: GetUserByLineID :one
SELECT * FROM users
WHERE line_id = @line_id::text;
