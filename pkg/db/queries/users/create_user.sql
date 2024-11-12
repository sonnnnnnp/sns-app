-- name: CreateUser :one
INSERT INTO users (
  line_id
) VALUES (
  sqlc.narg(line_id)::text
)
RETURNING users.id;
