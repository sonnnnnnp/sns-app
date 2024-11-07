-- name: DeleteUserFollower :exec
DELETE FROM user_followers
WHERE follower_id = $1 AND following_id = $2;
