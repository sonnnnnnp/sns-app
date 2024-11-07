-- name: DeletePostFavorite :exec
DELETE FROM post_favorites
WHERE user_id = $1 AND post_id = $2;
