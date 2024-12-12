package post

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/reverie/internal/http/adapter/api"
	internal_errors "github.com/sonnnnnnp/reverie/internal/http/errors"
	"github.com/sonnnnnnp/reverie/internal/http/tools/ctxhelper"
	"github.com/sonnnnnnp/reverie/pkg/db"
)

func (uc *PostUsecase) GetPostByID(ctx context.Context, pID uuid.UUID) (*api.Post, error) {
	selfUID := ctxhelper.GetUserID(ctx)

	r, err := db.New(uc.pool).GetPostByID(ctx, db.GetPostByIDParams{
		SelfID: selfUID,
		PostID: pID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal_errors.ErrPostNotFound
		}
		return nil, err
	}

	return &api.Post{
		Author: api.User{
			AvatarImageUrl: r.User.AvatarImageUrl,
			BannerImageUrl: r.User.BannerImageUrl,
			Biography:      r.User.Biography,
			CreatedAt:      r.User.CreatedAt.Time,
			Id:             r.User.ID,
			Name:           r.User.Name,
			Nickname:       r.User.Nickname,
			UpdatedAt:      r.User.UpdatedAt.Time,
		},
		Id:             r.Post.ID,
		Text:           r.Post.Text,
		Favorited:      r.Favorited,
		FavoritesCount: int(r.FavoritesCount),
		CreatedAt:      r.Post.CreatedAt.Time,
		UpdatedAt:      r.Post.UpdatedAt.Time,
	}, nil
}
