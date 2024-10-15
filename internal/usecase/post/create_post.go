package post

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/internal/tools/ws"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (pu *PostUsecase) CreatePost(ctx context.Context, body *oapi.CreatePostJSONBody) (*oapi.Post, error) {
	uID := ctxhelper.GetUserID(ctx)

	exists, err := pu.userRepo.GetUserExistence(ctx, uID)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.ErrUserNotFound
	}

	p, err := pu.postRepo.CreatePost(ctx, uID, body)
	if err != nil {
		return nil, err
	}

	// TODO: broadcast post to timeline channel
	websocket := ctxhelper.GetWebSocketHub(ctx)
	websocket.Broadcast <- ws.Message{Type: "post"}

	return &oapi.Post{
		Id: p.ID,
		Author: oapi.User{
			AvatarImageUrl: &p.Edges.Author.AvatarImageURL,
			Biography:      &p.Edges.Author.Biography,
			BannerImageUrl: &p.Edges.Author.BannerImageURL,
			CreatedAt:      p.Edges.Author.CreatedAt,
			Nickname:       p.Edges.Author.Nickname,
			Id:             p.Edges.Author.ID,
			UpdatedAt:      p.Edges.Author.UpdatedAt,
			Name:           p.Edges.Author.Name,
		},
		Text:      &p.Text,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}, nil
}
