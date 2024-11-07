-- name: GetUserFollowing :many
SELECT
    users.*,
    user_followers.created_at AS followed_at
FROM users
JOIN user_followers ON users.id = user_followers.following_id
WHERE user_followers.follower_id = @user_id
ORDER BY user_followers.created_at DESC;
