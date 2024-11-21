-- name: JoinCall :exec
INSERT INTO
  call_participants (
    call_id,
    participant_id,
    role
  )
VALUES
  (
    @call_id::uuid,
    @participant_id::uuid,
    sqlc.narg(role)::call_participant_role
  );
