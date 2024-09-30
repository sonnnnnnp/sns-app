package authorize

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

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

	isNew := false // 新規ユーザーを判定するためのフラグ

	u, err := au.userRepo.GetUserByLineID(ctx, lineUser.UserID)
	if err != nil {
		return nil, err
	}

	if u == nil {
		isNew = true

		u, err = au.userRepo.CreateUser(ctx)
		if err != nil {
			return nil, err
		}
	}

	return au.generateAuthorization([]byte(cfg.JWTSecret), u.ID, isNew)
}
