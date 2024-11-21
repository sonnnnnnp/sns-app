-- name: GetUserByName :one
SELECT
    sqlc.embed(users),
    (
        SELECT
            COUNT(*)
        FROM
            user_follows
        WHERE
            user_follows.follower_id = users.id
    ) AS following_count,
    (
        SELECT
            COUNT(*)
        FROM
            user_follows
        WHERE
            user_follows.followed_id = users.id
    ) AS followers_count,
    (
        SELECT
            COUNT(*)
        FROM
            posts
        WHERE
            posts.author_id = users.id
    ) AS posts_count,
    (
        SELECT
            COUNT(*)
        FROM
            posts
        WHERE
            -- TODO: メディアを含む投稿を探す
            posts.author_id = users.id
    ) AS media_count,
    (
        SELECT
            COUNT(*)
        FROM
            post_favorites
        WHERE
            post_favorites.user_id = users.id
    ) AS favorites_count
FROM
    users
WHERE
    name = @name::text;
