// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: get_self.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getSelf = `-- name: GetSelf :one
SELECT
    users.id, users.custom_id, users.nickname, users.biography, users.avatar_image_url, users.banner_image_url, users.is_private, users.birthdate, users.line_id, users.created_at, users.updated_at,
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
    id = $1::uuid
`

type GetSelfRow struct {
	User           User  `json:"user"`
	FollowingCount int64 `json:"following_count"`
	FollowersCount int64 `json:"followers_count"`
	PostsCount     int64 `json:"posts_count"`
	MediaCount     int64 `json:"media_count"`
	FavoritesCount int64 `json:"favorites_count"`
}

func (q *Queries) GetSelf(ctx context.Context, selfID uuid.UUID) (GetSelfRow, error) {
	row := q.db.QueryRow(ctx, getSelf, selfID)
	var i GetSelfRow
	err := row.Scan(
		&i.User.ID,
		&i.User.CustomID,
		&i.User.Nickname,
		&i.User.Biography,
		&i.User.AvatarImageUrl,
		&i.User.BannerImageUrl,
		&i.User.IsPrivate,
		&i.User.Birthdate,
		&i.User.LineID,
		&i.User.CreatedAt,
		&i.User.UpdatedAt,
		&i.FollowingCount,
		&i.FollowersCount,
		&i.PostsCount,
		&i.MediaCount,
		&i.FavoritesCount,
	)
	return i, err
}
