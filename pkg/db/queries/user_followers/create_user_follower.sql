-- name: CreateUserFollower :exec
INSERT INTO user_followers (
  follower_id, following_id
) VALUES (
  @follower_id::uuid, @following_id::uuid
);
