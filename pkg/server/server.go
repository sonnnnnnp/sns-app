package server

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"

	"github.com/sonnnnnnp/sns-app/internal"
	"github.com/sonnnnnnp/sns-app/internal/adapter/gateway/db"
	"github.com/sonnnnnnp/sns-app/internal/adapter/middleware"
	"github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ws"
	"github.com/sonnnnnnp/sns-app/pkg/config"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

func Run(cfg *config.Config) {
	db, err := db.Open(&db.ConnectionOptions{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Name:     cfg.DBName,
		Password: cfg.DBPassword,
		SSLMode:  cfg.DBSSLMode,
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

	middlewarecfg := echomiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}

	jwtExcludePaths := []string{
		"/authorize/line",
		"/authorize/refresh",
	}

	websocket := ws.NewHub()
	go websocket.Listen()

	swagger, err := oapi.GetSwagger()
	if err != nil {
		panic(err)
	}

	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.CORSWithConfig(middlewarecfg))
	e.Use(middleware.Config(cfg))
	e.Use(middleware.JWT(jwtExcludePaths))
	e.Use(middleware.WebSocket(websocket))
	e.Use(middleware.RequestValidator(swagger))

	oapi.RegisterHandlers(e, internal.Wire(db))

	e.Logger.Fatal(e.Start(":1323"))
}
