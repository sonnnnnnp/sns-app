-- name: DeleteUserFollow :exec
DELETE FROM user_follows
WHERE follower_id = @follower_id::uuid AND following_id = @following_id::uuid;
