package authorize

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *AuthorizeUsecase) createOrGetUserID(ctx context.Context, lineID string) (uID uuid.UUID, isNew bool, err error) {
	queries := db.New(uc.pool)

	u, err := queries.GetUserByLineID(ctx, lineID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			uID, err = queries.CreateUser(ctx, &lineID)
			if err != nil {
				return uuid.Nil, false, err
			}
			return uID, true, nil
		}
		return uuid.Nil, false, err
	}

	return u.ID, false, nil
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

	uID, isNew, err := uc.createOrGetUserID(ctx, lineUser.UserID)
	if err != nil {
		return nil, err
	}

	return uc.generateAuthorization([]byte(cfg.JWTSecret), uID, isNew)
}
