package authorize

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *AuthorizeUsecase) createOrGetUser(ctx context.Context, lineID string) (user *db.User, isNew bool, err error) {
	queries := db.New(uc.pool)

	var u db.User
	u, err = queries.GetUserByLineID(ctx, lineID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			u, err = queries.CreateUser(ctx, &lineID)
			if err != nil {
				return nil, false, err
			}
		}
		return nil, false, err
	}

	return &u, false, nil
}

func (uc *AuthorizeUsecase) AuthorizeWithLine(ctx context.Context, code string) (*api.Authorization, error) {
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
