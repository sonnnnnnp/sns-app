package server

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"

	"github.com/sonnnnnnp/sns-app/internal"
	"github.com/sonnnnnnp/sns-app/internal/adapter/gateway/db"
	"github.com/sonnnnnnp/sns-app/internal/adapter/middleware"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/pkg/config"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func Run(cfg *config.Config) {
	db, err := db.Open(&db.ConnectionOptions{
		Host:     "db",
		Port:     5432,
		User:     "user",
		Name:     "db",
		Password: "password",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	defer db.Close()

	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()

	e.HTTPErrorHandler = errors.ErrorHandler

	jwtExcludePaths := []string{
		"/authorize/line",
	}

	swagger, err := oapi.GetSwagger()
	if err != nil {
		panic(err)
	}

	e.Use(echomiddleware.Logger())
	e.Use(middleware.ConfigMiddleware(cfg))
	e.Use(middleware.JWTMiddleware(jwtExcludePaths))
	e.Use(middleware.RequestValidatorMiddleware(swagger))

	oapi.RegisterHandlers(e, internal.Wire(db))

	e.Logger.Fatal(e.Start(":1323"))
}
