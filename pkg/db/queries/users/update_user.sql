-- name: UpdateUser :exec
UPDATE users
SET
    name = COALESCE(sqlc.narg(name), name)::text,
    nickname = COALESCE(sqlc.narg(nickname), nickname)::text,
    biography = COALESCE(sqlc.narg(biography), biography)::text
WHERE users.id = @user_id::uuid;
