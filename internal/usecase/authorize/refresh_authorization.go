package authorize

import (
	"context"

	"github.com/sonnnnnnp/sns-app/internal/tools/ctxhelper"
	"github.com/sonnnnnnp/sns-app/internal/tools/jwter"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func (uc *AuthorizeUsecase) RefreshAuthorization(ctx context.Context, body *oapi.RefreshAuthorizationJSONBody) (*oapi.Authorization, error) {
	cfg := ctxhelper.GetConfig(ctx)

	uID, err := jwter.Verify(body.RefreshToken, "refresh", []byte(cfg.JWTSecret))
	if err != nil {
		return nil, err
	}

	return uc.generateAuthorization([]byte(cfg.JWTSecret), uID, false)
}
