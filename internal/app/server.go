package app

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/sonnnnnnp/sns-app/internal/pkg/oapi"
	"github.com/sonnnnnnp/sns-app/pkg/config"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) PostAuthorizeLine(ctx echo.Context) error {
	return ctx.JSON(http.StatusCreated, &oapi.LoginResponse{
		AccessToken:  "1234567890",
		RefreshToken: "abcdefghijklmnopqrstuvwxyz",
		Username:     "or4p90.fo0qg4",
	})
}

func (Server) GetUsersUserId(ctx echo.Context, userId string) error {
	return ctx.JSON(http.StatusOK, &oapi.User{
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
