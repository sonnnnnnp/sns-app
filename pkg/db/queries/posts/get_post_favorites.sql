-- name: GetPostFavorites :many
SELECT
    sqlc.embed(users),
    sqlc.embed(post_favorites)
FROM users
JOIN post_favorites ON users.id = post_favorites.user_id
WHERE post_favorites.post_id = $1
ORDER BY post_favorites.created_at DESC;
