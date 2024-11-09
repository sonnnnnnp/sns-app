-- name: DeleteUserBlock :exec
DELETE FROM user_blocks
WHERE blocker_id = @blocker_id::uuid AND blocking_id = @blocking_id::uuid;
