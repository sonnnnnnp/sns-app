-- name: CreateUserBlock :exec
INSERT INTO
  user_blocks (
    blocker_id,
    blocked_id
  )
VALUES
  (
    @blocker_id::uuid,
    @blocked_id::uuid
  );
