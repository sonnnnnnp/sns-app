-- name: GetPostFavorite :one
SELECT
    sqlc.embed(users),
    sqlc.embed(post_favorites)
FROM post_favorites
JOIN users ON post_favorites.user_id = users.id
WHERE post_favorites.post_id = $1 AND post_favorites.user_id = $2;
