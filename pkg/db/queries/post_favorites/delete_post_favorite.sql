-- name: DeletePostFavorite :exec
DELETE FROM post_favorites
WHERE user_id = @user_id::uuid AND post_id = @post_id::uuid;
