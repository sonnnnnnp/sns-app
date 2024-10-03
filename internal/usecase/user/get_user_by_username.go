package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uu *UserUsecase) GetUserByUsername(ctx context.Context, username string) (*ent.User, *oapi.SocialContext, error) {
	uID := ctxhelper.GetUserID(ctx)

	u, err := uu.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, nil, err
	}

	if u == nil {
		return nil, nil, nil
	}

	sc, err := uu.userRepo.GetSocialContext(ctx, uID, u)
	if err != nil {
		return nil, nil, err
	}

	return u, sc, nil
}
