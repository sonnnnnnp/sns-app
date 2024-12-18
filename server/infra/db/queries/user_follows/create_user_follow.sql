-- name: CreateUserFollow :exec
INSERT INTO
  user_follows (
    follower_id,
    followed_id
  )
VALUES
  (
    @follower_id::uuid,
    @followed_id::uuid
  );
