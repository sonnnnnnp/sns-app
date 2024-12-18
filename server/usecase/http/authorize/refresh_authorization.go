package authorize

import (
	"context"

	"github.com/sonnnnnnp/reverie/server/adapter/http/api"
	"github.com/sonnnnnnp/reverie/server/pkg/ctxhelper"
	"github.com/sonnnnnnp/reverie/server/pkg/jwter"
)

func (uc *AuthorizeUsecase) RefreshAuthorization(ctx context.Context, body *api.RefreshAuthorizationJSONBody) (*api.Authorization, error) {
	cfg := ctxhelper.GetConfig(ctx)

	uID, err := jwter.Verify(body.RefreshToken, "refresh", []byte(cfg.JWTSecret))
	if err != nil {
		return nil, err
	}

	return uc.generateAuthorization([]byte(cfg.JWTSecret), uID, false)
}
