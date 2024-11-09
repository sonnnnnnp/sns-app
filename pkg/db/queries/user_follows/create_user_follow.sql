-- name: CreateUserFollow :exec
INSERT INTO user_follows (
  follower_id, following_id
) VALUES (
  @follower_id::uuid, @following_id::uuid
);
