package post

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *PostUsecase) CreatePost(ctx context.Context, body *oapi.CreatePostJSONBody) (*oapi.Post, error) {
	uID := ctxhelper.GetUserID(ctx)

	pID, err := uc.postRepo.CreatePost(ctx, uID, body)
	if err != nil {
		return nil, err
	}

	r, err := uc.postRepo.GetPostByID(ctx, pID)
	if err != nil {
		return nil, err
	}

	return &oapi.Post{
		Author: oapi.User{
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
