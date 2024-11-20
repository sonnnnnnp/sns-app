-- name: CreateCall :one
INSERT INTO
  calls (
    host_id,
    title,
    type,
    joinable_by
  )
VALUES
  (
    @host_id::uuid,
    @title::text,
    @type::call_type,
    @joinable_by::call_joinable_by
  )
RETURNING calls.id;
