package authorize_usecase

import (
	"context"

	"github.com/google/uuid"
	user_repository "github.com/sonnnnnnp/sns-app/internal/domain/repository/user"
	"github.com/sonnnnnnp/sns-app/pkg/line"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IAuthorizeUsecase interface {
	AuthorizeWithLine(ctx context.Context, code string) (*oapi.Authorization, error)
}

type AuthorizeUsecase struct {
	line *line.Client

	userRepo *user_repository.UserRepository
}

func New(
	line *line.Client,
	userRepo *user_repository.UserRepository,
) *AuthorizeUsecase {
	return &AuthorizeUsecase{
		line: line,

		userRepo: userRepo,
	}
}

func (au *AuthorizeUsecase) generateAuthorization(uid uuid.UUID, IsNew bool) *oapi.Authorization {
	return &oapi.Authorization{
		AccessToken:  "AccessToken",
		RefreshToken: "RefreshToken",
		UserId:       uid.String(),
		IsNew:        IsNew,
	}
}

var _ IAuthorizeUsecase = (*AuthorizeUsecase)(nil)
