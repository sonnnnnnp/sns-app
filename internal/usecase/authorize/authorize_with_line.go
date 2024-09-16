package authorize_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (au *AuthorizeUsecase) AuthorizeWithLine(ctx context.Context, code string) (*oapi.Authorization, error) {
	resp, err := au.line.GetToken(
		code,
		au.cfg.LineAuthRedirectURL,
		au.cfg.LineAuthClientID,
		au.cfg.LineAuthClientSecret,
	)
	if err != nil {
		return nil, err
	}

	lineUser, err := au.line.GetUser(resp.AccessToken)
	if err != nil {
		return nil, err
	}

	u, err := au.userRepo.GetUserByLineID(ctx, lineUser.UserID)
	if err != nil {
		return nil, err
	}

	if u == nil {
		u, err = au.userRepo.CreateUser(ctx, uuid.NewString())
		if err != nil {
			return nil, err
		}
	}

	return &oapi.Authorization{
		AccessToken:  lineUser.UserID,
		RefreshToken: "RefreshToken",
		UserId:       u.ID.String(),
	}, nil
}
