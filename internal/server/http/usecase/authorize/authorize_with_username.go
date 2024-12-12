package authorize

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/sonnnnnnp/reverie/internal/pkg/ctxhelper"
	internal_errors "github.com/sonnnnnnp/reverie/internal/pkg/errors"
	"github.com/sonnnnnnp/reverie/internal/server/http/adapter/api"
	"github.com/sonnnnnnp/reverie/pkg/db"
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
