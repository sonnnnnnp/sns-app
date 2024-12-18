-- name: LeaveCall :exec
UPDATE
    call_participants
SET
    left_at = NOW()
WHERE
    call_participants.call_id = @call_id::uuid
    AND call_participants.participant_id = @participant_id::uuid;
