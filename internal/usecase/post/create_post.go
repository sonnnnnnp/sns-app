package post

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (pu *PostUsecase) CreatePost(ctx context.Context, body *oapi.CreatePostJSONBody) (*ent.Post, error) {
	uID := ctxhelper.GetUserID(ctx)

	u, err := pu.userRepo.GetUserByID(ctx, uID)
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, errors.ErrUserNotFound
	}

	p, err := pu.postRepo.CreatePost(ctx, uID, body)
	if err != nil {
		return nil, err
	}

	return p, nil
}
