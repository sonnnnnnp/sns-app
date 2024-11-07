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
    users.id, users.name, users.nickname, users.biography, users.avatar_image_url, users.banner_image_url, users.birthdate, users.line_id, users.created_at, users.updated_at,
    post_favorites.user_id, post_favorites.post_id, post_favorites.created_at
FROM post_favorites
JOIN users ON post_favorites.user_id = users.id
WHERE post_favorites.post_id = $1 AND post_favorites.user_id = $2
`

type GetPostFavoriteParams struct {
	PostID uuid.UUID
	UserID uuid.UUID
}

type GetPostFavoriteRow struct {
	User         User
	PostFavorite PostFavorite
}

func (q *Queries) GetPostFavorite(ctx context.Context, arg GetPostFavoriteParams) (GetPostFavoriteRow, error) {
	row := q.db.QueryRow(ctx, getPostFavorite, arg.PostID, arg.UserID)
	var i GetPostFavoriteRow
	err := row.Scan(
		&i.User.ID,
		&i.User.Name,
		&i.User.Nickname,
		&i.User.Biography,
		&i.User.AvatarImageUrl,
		&i.User.BannerImageUrl,
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
