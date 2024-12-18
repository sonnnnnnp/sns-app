package post

import (
	"context"

	"github.com/sonnnnnnp/reverie/server/adapter/http/api"
	"github.com/sonnnnnnp/reverie/server/infra/db"
	"github.com/sonnnnnnp/reverie/server/pkg/ctxhelper"
)

func (uc *PostUsecase) CreatePost(ctx context.Context, body *api.CreatePostJSONBody) (*api.Post, error) {
	queries := db.New(uc.pool)

	selfUID := ctxhelper.GetUserID(ctx)

	pID, err := queries.CreatePost(ctx, db.CreatePostParams{
		AuthorID:  selfUID,
		Text:      body.Text,
		ReplyToID: body.ReplyToPostId,
		RepostID:  body.RepostPostId,
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
