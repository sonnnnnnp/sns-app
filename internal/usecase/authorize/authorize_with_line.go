package authorize

import (
	"context"
	"errors"

	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"

	internal_errors "github.com/sonnnnnnp/sns-app/internal/errors"
)

func (au *AuthorizeUsecase) createOrGetUser(ctx context.Context, lineID string) (u *ent.User, isNew bool, err error) {
	u, err = au.userRepo.GetUserByLineID(ctx, lineID)
	if err != nil {
		if errors.Is(err, internal_errors.ErrUserNotFound) {
			u, err = au.userRepo.CreateUser(ctx, &lineID)
			if err != nil {
				return nil, false, err
			}
			return u, true, nil
		}
		return nil, false, err
	}
	return u, false, nil
}

func (au *AuthorizeUsecase) AuthorizeWithLine(ctx context.Context, code string) (*oapi.Authorization, error) {
	cfg := ctxhelper.GetConfig(ctx)

	resp, err := au.line.GetToken(
		code,
		cfg.LineAuthRedirectURL,
		cfg.LineAuthClientID,
		cfg.LineAuthClientSecret,
	)
	if err != nil {
		return nil, err
	}

	lineUser, err := au.line.GetUser(resp.AccessToken)
	if err != nil {
		return nil, err
	}

	u, isNew, err := au.createOrGetUser(ctx, lineUser.UserID)
	if err != nil {
		return nil, err
	}

	return au.generateAuthorization([]byte(cfg.JWTSecret), u.ID, isNew)
}
