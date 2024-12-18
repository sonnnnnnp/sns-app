// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: get_post_favorite.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getPostFavorite = `-- name: GetPostFavorite :one
SELECT
    users.id, users.custom_id, users.nickname, users.biography, users.avatar_image_url, users.banner_image_url, users.is_private, users.birthdate, users.line_id, users.created_at, users.updated_at,
    post_favorites.user_id, post_favorites.post_id, post_favorites.created_at
FROM
    post_favorites
    INNER JOIN
        users
        ON post_favorites.user_id = users.id
WHERE
    post_favorites.post_id = $1::uuid
    AND post_favorites.user_id = $2::uuid
`

type GetPostFavoriteParams struct {
	PostID uuid.UUID `json:"post_id"`
	UserID uuid.UUID `json:"user_id"`
}

type GetPostFavoriteRow struct {
	User         User         `json:"user"`
	PostFavorite PostFavorite `json:"post_favorite"`
}

func (q *Queries) GetPostFavorite(ctx context.Context, arg GetPostFavoriteParams) (GetPostFavoriteRow, error) {
	row := q.db.QueryRow(ctx, getPostFavorite, arg.PostID, arg.UserID)
	var i GetPostFavoriteRow
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
		&i.PostFavorite.UserID,
		&i.PostFavorite.PostID,
		&i.PostFavorite.CreatedAt,
	)
	return i, err
}
