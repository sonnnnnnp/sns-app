package post

import (
	"context"

	"github.com/sonnnnnnp/reverie/internal/adapter/api"
	"github.com/sonnnnnnp/reverie/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/reverie/pkg/db"
)

func (uc *PostUsecase) CreatePost(ctx context.Context, body *api.CreatePostJSONBody) (*api.Post, error) {
	queries := db.New(uc.pool)

	selfUID := ctxhelper.GetUserID(ctx)

	pID, err := queries.CreatePost(ctx, db.CreatePostParams{
		AuthorID:  selfUID,
		Text:      body.Content,
		ReplyToID: body.ReplyToPostId,
	})
	if err != nil {
		return nil, err
	}

	r, err := queries.GetPostByID(ctx, db.GetPostByIDParams{
		SelfID: selfUID,
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
