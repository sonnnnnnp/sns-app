-- name: GetUserByID :one
SELECT * FROM users
WHERE id = @user_id::uuid;
