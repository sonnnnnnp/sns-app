-- name: GetUserCallTimeline :many
SELECT
    sqlc.embed(calls),
    json_agg(json_build_object(
        'role', call_participants.role,
        'user', to_jsonb(users) || jsonb_build_object(
            'block_status', jsonb_build_object(
                'blocking', EXISTS (
                    SELECT 1
                    FROM user_blocks
                    WHERE blocker_id = @self_id::uuid
                    AND blocked_id = users.id
                ),
                'blocked_by', EXISTS (
                    SELECT 1
                    FROM user_blocks
                    WHERE blocker_id = users.id
                    AND blocked_id = @self_id::uuid
                )
            )
        )
    )) AS participants
FROM
    calls
    LEFT JOIN
        call_participants
        ON calls.id = call_participants.call_id
    LEFT JOIN
        users
        ON call_participants.participant_id = users.id
WHERE
    calls.host_id = @user_id::uuid
GROUP BY
    calls.id
ORDER BY
    RANDOM()
LIMIT
    $1;
