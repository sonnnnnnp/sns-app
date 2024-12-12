package authorize

import (
	"context"

	"github.com/sonnnnnnp/reverie/internal/http/adapter/api"
	"github.com/sonnnnnnp/reverie/internal/http/tools/ctxhelper"
	"github.com/sonnnnnnp/reverie/internal/http/tools/jwter"
)

func (uc *AuthorizeUsecase) RefreshAuthorization(ctx context.Context, body *api.RefreshAuthorizationJSONBody) (*api.Authorization, error) {
	cfg := ctxhelper.GetConfig(ctx)

	uID, err := jwter.Verify(body.RefreshToken, "refresh", []byte(cfg.JWTSecret))
	if err != nil {
		return nil, err
	}

	return uc.generateAuthorization([]byte(cfg.JWTSecret), uID, false)
}
