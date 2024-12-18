-- name: EndCall :exec
UPDATE
    calls
SET
    ended_at = NOW()
WHERE
    calls.id = @call_id::uuid;
