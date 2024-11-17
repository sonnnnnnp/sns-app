package post

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *PostUsecase) CreatePostReply(ctx context.Context, pID uuid.UUID, body *api.CreatePostReplyJSONBody) (*api.Post, error) {
	queries := db.New(uc.pool)

	selfUID := ctxhelper.GetUserID(ctx)

	pID, err := queries.CreatePostReply(ctx, db.CreatePostReplyParams{
		AuthorID:  selfUID,
		ReplyToID: pID,
		Text:      body.Content,
	})
	if err != nil {
		return nil, err
	}

	r, err := queries.GetPostByID(ctx, db.GetPostByIDParams{
		SelfID: &selfUID,
		PostID: pID,
	})
	if err != nil {
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
		Favorited:      false,
		FavoritesCount: 0,
		CreatedAt:      r.Post.CreatedAt.Time,
		UpdatedAt:      r.Post.UpdatedAt.Time,
	}, nil
}
