-- name: GetPostByID :one
SELECT
    sqlc.embed(posts),
    sqlc.embed(users),
    (
        SELECT
            COUNT(*)
        FROM
            post_favorites
        WHERE
            post_favorites.post_id = posts.id
    ) AS favorites_count,
    EXISTS (
        SELECT
            1
        FROM
            post_favorites
        WHERE
            post_favorites.post_id = posts.id
            AND (
                post_favorites.user_id = @self_id::uuid
            )
    ) AS favorited
FROM
    posts
    INNER JOIN
        users
        ON posts.author_id = users.id
WHERE
    posts.id = @post_id::uuid;
