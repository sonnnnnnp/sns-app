-- name: HasUserActiveCall :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            call_participants
        WHERE
            call_participants.participant_id = @user_id::uuid
    ) AS joined;
