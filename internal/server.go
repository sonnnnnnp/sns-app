package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"

	"github.com/sonnnnnnp/sns-app/pkg/config"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type Response struct {
	Ok   bool        `json:"ok"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Server struct {
	db *ent.Client
}

func NewServer(db *ent.Client) Server {
	return Server{
		db: db,
	}
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
	u, err := s.db.Debug().User.Create().SetUsername(userId).SetDisplayName(userId).Save(ctx.Request().Context())
	if err != nil {
		return s.json(ctx, &struct {
			Message string `json:"message"`
		}{Message: fmt.Sprintf("%e", err)})
	}

	return s.json(ctx, &oapi.User{
		Id:          u.ID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		AvatarUrl:   u.AvatarURL,
		CoverUrl:    u.CoverURL,
		Biography:   u.Biography,
		Birthdate:   u.Birthdate,
		UpdatedAt:   u.UpdatedAt,
		CreatedAt:   u.CreatedAt,
	})
}

func Init(cfg *config.Config) {
	db, err := ent.Open("postgres", "host=db port=5432 user=user dbname=db password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer db.Close()

	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	server := NewServer(db)

	e := echo.New()

	oapi.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":1323"))
}
