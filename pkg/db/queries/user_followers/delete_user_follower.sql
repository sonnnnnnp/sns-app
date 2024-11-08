-- name: DeleteUserFollower :exec
DELETE FROM user_followers
WHERE follower_id = @follower_id::uuid AND following_id = @following_id::uuid;
