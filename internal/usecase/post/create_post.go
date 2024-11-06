package post

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *PostUsecase) CreatePost(ctx context.Context, body *oapi.CreatePostJSONBody) (*oapi.Post, error) {
	uID := ctxhelper.GetUserID(ctx)

	p, err := uc.postRepo.CreatePost(ctx, uID, body)
	if err != nil {
		return nil, err
	}

	return &oapi.Post{
		Id:        p.ID,
		Text:      p.Text,
		CreatedAt: p.CreatedAt.Time,
		UpdatedAt: p.UpdatedAt.Time,
	}, nil

	// exists, err := pu.userRepo.GetUserExistence(ctx, uID)
	// if err != nil {
	// 	return nil, err
	// }

	// if !exists {
	// 	return nil, errors.ErrUserNotFound
	// }

	// p, err := pu.postRepo.CreatePost(ctx, uID, body)
	// if err != nil {
	// 	return nil, err
	// }

	// data := &oapi.Post{
	// 	Id: p.ID,
	// 	Author: oapi.User{
	// 		AvatarImageUrl: &p.Edges.Author.AvatarImageURL,
	// 		Biography:      &p.Edges.Author.Biography,
	// 		BannerImageUrl: &p.Edges.Author.BannerImageURL,
	// 		CreatedAt:      p.Edges.Author.CreatedAt,
	// 		Nickname:       p.Edges.Author.Nickname,
	// 		Id:             p.Edges.Author.ID,
	// 		UpdatedAt:      p.Edges.Author.UpdatedAt,
	// 		Name:           p.Edges.Author.Name,
	// 	},
	// 	Text:      &p.Text,
	// 	CreatedAt: p.CreatedAt,
	// 	UpdatedAt: p.UpdatedAt,
	// }

	// pu.streamUsecase.Broadcast(
	// 	ctx,
	// 	ws.ChannelTimeline,
	// 	ws.Message{
	// 		Type: "post",
	// 		Body: data,
	// 	},
	// 	nil, // broadcast to all users
	// )

	// return data, nil
}
