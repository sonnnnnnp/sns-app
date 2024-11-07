package authorize

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	user_repository "github.com/sonnnnnnp/sns-app/internal/domain/user"
	"github.com/sonnnnnnp/sns-app/pkg/line"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type IAuthorizeUsecase interface {
	AuthorizeWithLine(ctx context.Context, code string) (*oapi.Authorization, error)
	RefreshAuthorization(ctx context.Context, body *oapi.RefreshAuthorizationJSONBody) (*oapi.Authorization, error)
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

var _ IAuthorizeUsecase = (*AuthorizeUsecase)(nil)

func (uc *AuthorizeUsecase) generateToken(jwtSecret []byte, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (uc *AuthorizeUsecase) generateAuthorization(jwtSecret []byte, uid uuid.UUID, IsNew bool) (*oapi.Authorization, error) {
	atoken, err := uc.generateToken(
		jwtSecret,
		jwt.MapClaims{
			"sub":   uid.String(),
			"exp":   jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			"scope": "access",
		},
	)
	if err != nil {
		return nil, err
	}

	rtoken, err := uc.generateToken(
		jwtSecret,
		jwt.MapClaims{
			"sub":   uid.String(),
			"exp":   jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			"scope": "refresh",
		},
	)
	if err != nil {
		return nil, err
	}

	return &oapi.Authorization{
		AccessToken:  atoken,
		RefreshToken: rtoken,
		UserId:       uid.String(),
		IsNew:        IsNew,
	}, nil
}
