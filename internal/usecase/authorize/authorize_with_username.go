package authorize

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/sns-app/internal/adapter/api"
	internal_errors "github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/pkg/db"
)

func (uc *AuthorizeUsecase) AuthorizeWithUsername(ctx context.Context, name string) (*api.Authorization, error) {
	cfg := ctxhelper.GetConfig(ctx)

	if cfg.APPEnv != "development" {
		return nil, internal_errors.ErrUnauthorized
	}

	queries := db.New(uc.pool)

	row, err := queries.GetUserByName(ctx, name)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, internal_errors.ErrUserNotFound
		}
		return nil, err
	}

	return uc.generateAuthorization([]byte(cfg.JWTSecret), row.User.ID, false)
}
