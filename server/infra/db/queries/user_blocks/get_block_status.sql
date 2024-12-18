-- name: GetBlockStatus :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            user_blocks
        WHERE
            blocker_id = @self_id::uuid
            AND blocked_id = @target_id::uuid
    ) AS blocking,
    EXISTS (
        SELECT
            1
        FROM
            user_blocks
        WHERE
            blocker_id = @target_id::uuid
            AND blocked_id = @self_id::uuid
    ) AS blocked_by;
