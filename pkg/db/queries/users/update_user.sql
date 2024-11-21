-- name: UpdateUser :exec
UPDATE
    users
SET
    name = COALESCE(sqlc.narg(name), name)::text,
    nickname = COALESCE(sqlc.narg(nickname), nickname)::text,
    biography = COALESCE(sqlc.narg(biography), biography)::text,
    avatar_image_url = COALESCE(sqlc.narg(avatar_image_url), avatar_image_url)::text,
    banner_image_url = COALESCE(sqlc.narg(banner_image_url), banner_image_url)::text,
    birthdate = COALESCE(sqlc.narg(birthdate), birthdate)::timestamptz,
    line_id = COALESCE(sqlc.narg(line_id), line_id)::text,
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at)::timestamptz
WHERE
    users.id = @user_id::uuid;
