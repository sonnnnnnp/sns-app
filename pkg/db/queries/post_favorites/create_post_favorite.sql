-- name: CreatePostFavorite :exec
INSERT INTO post_favorites (
  user_id, post_id
) VALUES (
  $1, $2
);
