-- name: DeletePost :exec
DELETE FROM posts
WHERE id = @post_id::uuid;
