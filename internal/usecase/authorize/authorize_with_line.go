package authorize

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *AuthorizeUsecase) createOrGetUser(ctx context.Context, lineID string) (u *db.User, isNew bool, err error) {
	u, err = uc.userRepo.GetUserByLineID(ctx, lineID)
	if err != nil {
		return nil, false, err
	}

	if u == nil {
		u, err = uc.userRepo.CreateUser(ctx, &lineID)
		if err != nil {
			return nil, false, err
		}
		return u, true, nil
	}

	return u, false, nil
}

func (uc *AuthorizeUsecase) AuthorizeWithLine(ctx context.Context, code string) (*oapi.Authorization, error) {
	cfg := ctxhelper.GetConfig(ctx)

	resp, err := uc.line.GetToken(
		code,
		cfg.LineAuthRedirectURL,
		cfg.LineAuthClientID,
		cfg.LineAuthClientSecret,
	)
	if err != nil {
		return nil, err
	}

	lineUser, err := uc.line.GetUser(resp.AccessToken)
	if err != nil {
		return nil, err
	}

	u, isNew, err := uc.createOrGetUser(ctx, lineUser.UserID)
	if err != nil {
		return nil, err
	}

	return uc.generateAuthorization([]byte(cfg.JWTSecret), u.ID, isNew)
}
