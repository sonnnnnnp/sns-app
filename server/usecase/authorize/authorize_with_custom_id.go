package authorize

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/reverie/server/adapter/http/api"
	"github.com/sonnnnnnp/reverie/server/infra/db"
	"github.com/sonnnnnnp/reverie/server/pkg/ctxhelper"
	internal_errors "github.com/sonnnnnnp/reverie/server/pkg/errors"
)

func (uc *AuthorizeUsecase) AuthorizeWithCustomID(ctx context.Context, customID string) (*api.Authorization, error) {
	cfg := ctxhelper.GetConfig(ctx)

	if cfg.APPEnv != "development" {
		return nil, internal_errors.ErrUnauthorized
	}

	queries := db.New(uc.pool)

	row, err := queries.GetUserByCustomID(ctx, customID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal_errors.ErrUserNotFound
		}
		return nil, err
	}

	return uc.generateAuthorization([]byte(cfg.JWTSecret), row.User.ID, false)
}
