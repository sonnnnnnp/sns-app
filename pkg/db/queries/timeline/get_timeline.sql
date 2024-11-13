-- name: GetTimeline :many
SELECT
    sqlc.embed(posts),
    sqlc.embed(users),
    (
        SELECT COUNT(*)
        FROM post_favorites
        WHERE post_favorites.post_id = posts.id
    ) AS favorites_count,
    EXISTS (
        SELECT 1
        FROM post_favorites
        WHERE post_favorites.post_id = posts.id AND (
            post_favorites.user_id = sqlc.narg(self_id)::uuid
        )
    ) AS favorited
FROM
    posts
INNER JOIN
    users ON posts.author_id = users.id
LEFT JOIN
    user_follows ON users.id = user_follows.following_id
LEFT JOIN
    user_blocks AS blocker ON
    users.id = blocker.blocking_id
    AND blocker.blocker_id = sqlc.narg(self_id)::uuid
LEFT JOIN
    user_blocks AS blocking ON
    users.id = blocking.blocker_id
    AND blocking.blocking_id = sqlc.narg(self_id)::uuid
WHERE
    (
        sqlc.narg(created_at)::timestamptz IS NULL
        OR posts.created_at < sqlc.narg(created_at)::timestamptz
    )
    AND (
        NOT @only_following::boolean
        OR user_follows.follower_id = sqlc.narg(self_id)::uuid
        OR posts.author_id = sqlc.narg(self_id)::uuid
    )
    AND blocker.blocking_id IS NULL
    AND blocking.blocker_id IS NULL
ORDER BY
    posts.created_at DESC
LIMIT $1;
