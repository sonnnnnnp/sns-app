package authorize_usecase

import (
	"context"

	user_repository "github.com/sonnnnnnp/sns-app/internal/domain/repository/user"
	"github.com/sonnnnnnp/sns-app/pkg/config"
	"github.com/sonnnnnnp/sns-app/pkg/line"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IAuthorizeUsecase interface {
	AuthorizeWithLine(ctx context.Context, code string) (*oapi.Authorization, error)
}

type AuthorizeUsecase struct {
	cfg  *config.Config
	line *line.Client

	userRepo *user_repository.UserRepository
}

func New(
	cfg *config.Config,
	line *line.Client,
	userRepo *user_repository.UserRepository,
) *AuthorizeUsecase {
	return &AuthorizeUsecase{
		cfg:  cfg,
		line: line,

		userRepo: userRepo,
	}
}

var _ IAuthorizeUsecase = (*AuthorizeUsecase)(nil)
