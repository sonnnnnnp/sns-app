package app

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sonnnnnnp/sns-app/internal/pkg/oapi"
	"github.com/sonnnnnnp/sns-app/pkg/config"
)

type Response struct {
	Ok   bool        `json:"ok"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s Server) json(ctx echo.Context, data interface{}) error {
	return ctx.JSON(http.StatusOK, &Response{
		Code: http.StatusOK,
		Data: data,
		Ok:   true,
	})
}

func (s Server) AuthorizeWithLINE(ctx echo.Context) error {
	return s.json(ctx, &oapi.Authorization{
		AccessToken:  "1234567890",
		RefreshToken: "abcdefghijklmnopqrstuvwxyz",
		UserId:       "or4p90.fo0qg4",
	})
}

func (s Server) RefreshAuthorization(ctx echo.Context) error {
	return s.json(ctx, &oapi.Authorization{
		AccessToken:  "1234567890",
		RefreshToken: "abcdefghijklmnopqrstuvwxyz",
		UserId:       "or4p90.fo0qg4",
	})
}

func (s Server) GetUser(ctx echo.Context, userId string) error {
	return s.json(ctx, &oapi.User{
		Id:          userId,
		Username:    "or4p90.fo0qg4",
		DisplayName: "辰男_伊藤31",
		AvatarUrl:   nil,
		CoverUrl:    nil,
		Biography:   nil,
		Birthdate:   nil,
		UpdatedAt:   "2024-02-03T12:28:47.053Z",
		CreatedAt:   "2024-02-03T12:28:47.053Z",
	})
}

func Init(cfg *config.Config) {
	server := NewServer()

	e := echo.New()

	oapi.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":1323"))
}
