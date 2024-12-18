-- name: CreatePost :one
INSERT INTO
  posts (
    author_id,
    text,
    reply_to_id,
    repost_id
  )
VALUES
  (
    @author_id::uuid,
    sqlc.narg(text)::text,
    sqlc.narg(reply_to_id)::uuid,
    sqlc.narg(repost_id)::uuid
  )
RETURNING posts.id;
