-- name: CreateUserBlock :exec
INSERT INTO user_blocks (
  blocker_id, blocking_id
) VALUES (
  @blocker_id::uuid, @blocking_id::uuid
);
