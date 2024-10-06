package user

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uu *UserUsecase) GetUserByName(ctx context.Context, name string) (*ent.User, *oapi.SocialContext, error) {
	uID := ctxhelper.GetUserID(ctx)

	u, err := uu.userRepo.GetUserByName(ctx, name)
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
