-- name: CreatePostFavorite :exec
INSERT INTO post_favorites (
  user_id, post_id
) VALUES (
  @user_id::uuid, @post_id::uuid
);
