-- name: GetPosts :many
SELECT
    sqlc.embed(posts),
    sqlc.embed(users),
    (
        SELECT COUNT(*) FROM post_favorites
        WHERE post_favorites.post_id = posts.id
    ) AS favorites_count,
    EXISTS (
        SELECT 1 FROM post_favorites
        WHERE post_favorites.post_id = posts.id AND post_favorites.user_id = $1
    ) AS favorited
FROM posts
JOIN users ON posts.author_id = users.id
ORDER BY posts.created_at DESC;
