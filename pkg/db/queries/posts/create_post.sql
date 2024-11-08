-- name: CreatePost :one
INSERT INTO posts (
  author_id, text
) VALUES (
  @author_id::uuid, sqlc.narg(text)::text
)
RETURNING posts.id;
