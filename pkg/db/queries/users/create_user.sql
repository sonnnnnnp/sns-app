-- name: CreateUser :one
INSERT INTO users (
  id, name, nickname, line_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;
