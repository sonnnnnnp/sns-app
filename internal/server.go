package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"

	"github.com/sonnnnnnp/sns-app/internal/middleware"
	"github.com/sonnnnnnp/sns-app/pkg/config"
	"github.com/sonnnnnnp/sns-app/pkg/ent"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type Response struct {
	OK   bool        `json:"ok"`
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

func (s Server) json(ctx echo.Context, code int, data interface{}) error {
	return ctx.JSON(http.StatusOK, &Response{
		Code: code,
		OK:   code == http.StatusOK,
		Data: data,
	})
}

func (s Server) AuthorizeWithLine(ctx echo.Context, params oapi.AuthorizeWithLineParams) error {
	return s.json(ctx, http.StatusOK, &oapi.Authorization{
		AccessToken:  "1234567890",
		RefreshToken: "abcdefghijklmnopqrstuvwxyz",
		UserId:       params.Code,
	})
}

func (s Server) RefreshAuthorization(ctx echo.Context) error {
	return s.json(ctx, http.StatusOK, &oapi.Authorization{
		AccessToken:  "1234567890",
		RefreshToken: "abcdefghijklmnopqrstuvwxyz",
		UserId:       "or4p90.fo0qg4",
	})
}

func (s Server) GetUser(ctx echo.Context, userId string) error {
	u, err := s.db.User.Create().SetUsername(userId).SetDisplayName(userId).Save(ctx.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "failed to create a user")
	}

	return s.json(ctx, http.StatusOK, &oapi.User{
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

	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			ctx.JSON(http.StatusOK, &Response{
				Code: he.Code,
				OK:   false,
				Data: he.Message,
			})
			return
		}
		ctx.JSON(http.StatusOK, &Response{
			Code: http.StatusInternalServerError,
			OK:   false,
			Data: &struct {
				Message string `json:"message"`
			}{
				// TODO: hide error details under production mode
				Message: fmt.Sprintf("Internal server error: %v", err),
			},
		})
	}

	swagger, err := oapi.GetSwagger()
	if err != nil {
		panic(err)
	}

	e.Use(echomiddleware.Logger())
	e.Use(middleware.RequestValidatorMiddleware(swagger))

	oapi.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":1323"))
}
