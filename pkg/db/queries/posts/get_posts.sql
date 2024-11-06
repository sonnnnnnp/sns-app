-- name: GetPosts :many
SELECT
    sqlc.embed(posts),
    sqlc.embed(users)
FROM posts
INNER JOIN users ON posts.author_id = users.id
ORDER BY posts.created_at DESC;
