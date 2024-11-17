-- name: CreatePostReply :one
INSERT INTO posts (
    author_id, reply_to_id, text
) VALUES (
    @author_id::uuid, @reply_to_id::uuid, sqlc.narg(text)::text
)
RETURNING posts.id;
