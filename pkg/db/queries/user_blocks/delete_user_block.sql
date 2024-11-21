-- name: DeleteUserBlock :exec
DELETE FROM
    user_blocks
WHERE
    blocker_id = @blocker_id::uuid
    AND blocked_id = @blocked_id::uuid;
