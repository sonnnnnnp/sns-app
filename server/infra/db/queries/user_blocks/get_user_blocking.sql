-- name: GetUserBlocking :many
SELECT
    users.*,
    user_blocks.created_at AS blocked_at
FROM
    users
    INNER JOIN
        user_blocks
        ON users.id = user_blocks.blocked_id
WHERE
    user_blocks.blocker_id = @user_id::uuid
ORDER BY
    user_blocks.created_at DESC;
