-- name: DeleteUserFollow :exec
DELETE FROM
    user_follows
WHERE
    follower_id = @follower_id::uuid
    AND followed_id = @followed_id::uuid;
