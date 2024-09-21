package post_usecase

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/domain/ent"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (pu *PostUsecase) CreatePost(ctx context.Context, body *oapi.CreatePostJSONBody) (*ent.Post, *ent.User, error) {
	uID := ctxhelper.GetUserID(ctx)

	p, u, err := pu.postRepo.CreatePost(ctx, uID, body)
	if err != nil {
		return nil, nil, err
	}

	return p, u, nil
}
