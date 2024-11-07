-- name: CreateUserFollower :exec
INSERT INTO user_followers (
  follower_id, following_id
) VALUES (
  $1, $2
);
