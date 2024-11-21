-- name: IsUserCallHost :one
SELECT
    EXISTS (
        SELECT
            1
        FROM
            calls
        WHERE
            calls.host_id = @user_id::uuid
            AND calls.id = @call_id::uuid
    ) AS is_host;
