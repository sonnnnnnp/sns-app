package gateway

import (
	"context"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sonnnnnnp/reverie/internal/pkg/ctxhelper"
	"github.com/sonnnnnnp/reverie/internal/server/http/adapter/api"
)

func (uc *GatewayUsecase) GetGatewayURI(ctx context.Context) (*api.GatewayURI, error) {
	cfg := ctxhelper.GetConfig(ctx)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   ctxhelper.GetUserID(ctx).String(),
		"exp":   jwt.NewNumericDate(time.Now().Add(30 * time.Second)), // WebSocket では認証ヘッダーを利用できないため使い捨てトークンを生成する
		"scope": "gateway",
	})

	signedToken, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return nil, err
	}

	url, err := url.Parse(cfg.WSURL)
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Set("t", signedToken)
	url.RawQuery = q.Encode()

	return &api.GatewayURI{
		Uri: url.String(),
	}, nil
}
