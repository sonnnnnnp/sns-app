-- name: GetPostFavorite :one
SELECT
    sqlc.embed(users),
    sqlc.embed(post_favorites)
FROM post_favorites
INNER JOIN users ON post_favorites.user_id = users.id
WHERE post_favorites.post_id = @post_id::uuid AND post_favorites.user_id = @user_id::uuid;
