-- name: GetSocialConnection :one
SELECT
    EXISTS (
        SELECT 1
        FROM user_follows
        WHERE follower_id = @self_id::uuid AND following_id = @target_id::uuid
    ) AS following,
    EXISTS (
        SELECT 1
        FROM user_follows
        WHERE follower_id = @target_id::uuid AND following_id = @self_id::uuid
    ) AS followed_by;
