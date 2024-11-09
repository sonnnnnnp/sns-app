// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: update_user.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET
    name = COALESCE($1, name)::text,
    nickname = COALESCE($2, nickname)::text,
    biography = COALESCE($3, biography)::text,
    avatar_image_url = COALESCE($4, avatar_image_url)::text,
    banner_image_url = COALESCE($5, banner_image_url)::text,
    birthdate = COALESCE($6, birthdate)::timestamptz,
    line_id = COALESCE($7, line_id)::text,
    updated_at = COALESCE($8, updated_at)::timestamptz
WHERE users.id = $9::uuid
`

type UpdateUserParams struct {
	Name           *string
	Nickname       *string
	Biography      *string
	AvatarImageUrl *string
	BannerImageUrl *string
	Birthdate      *time.Time
	LineID         *string
	UpdatedAt      *time.Time
	UserID         uuid.UUID
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.Name,
		arg.Nickname,
		arg.Biography,
		arg.AvatarImageUrl,
		arg.BannerImageUrl,
		arg.Birthdate,
		arg.LineID,
		arg.UpdatedAt,
		arg.UserID,
	)
	return err
}