-- name: GetUserByID :one
SELECT
    sqlc.embed(users),
    (
        SELECT
            COUNT(*)
        FROM
            user_follows
        WHERE
            user_follows.follower_id = users.id
    ) as following_count,
    (
        SELECT
            COUNT(*)
        FROM
            user_follows
        WHERE
            user_follows.following_id = users.id
    ) as followers_count,
    (
        SELECT
            COUNT(*)
        FROM
            posts
        WHERE
            posts.author_id = users.id
    ) as posts_count,
    (
        SELECT
            COUNT(*)
        FROM
            posts
        WHERE
            -- TODO: メディアを含む投稿を探す
            posts.author_id = users.id
    ) as media_count,
    (
        SELECT
            COUNT(*)
        FROM
            post_favorites
        WHERE
            post_favorites.user_id = users.id
    ) as favorites_count
FROM
    users
WHERE
    id = @user_id::uuid;
