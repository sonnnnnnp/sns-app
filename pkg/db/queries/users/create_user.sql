-- name: CreateUser :one
INSERT INTO users (
  line_id
) VALUES (
  $1
)
RETURNING *;
