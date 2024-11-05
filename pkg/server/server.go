package server

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/sonnnnnnp/sns-app/internal"
	"github.com/sonnnnnnp/sns-app/internal/adapter/middleware"
	internal_errors "github.com/sonnnnnnp/sns-app/internal/errors"
	"github.com/sonnnnnnp/sns-app/internal/tools/ws"
	"github.com/sonnnnnnp/sns-app/pkg/config"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Run(cfg *config.Config) {
	// sql, err := db.Open(&db.ConnectionOptions{
	// 	Host:     cfg.DBHost,
	// 	Port:     cfg.DBPort,
	// 	User:     cfg.DBUser,
	// 	Name:     cfg.DBName,
	// 	Password: cfg.DBPassword,
	// 	SSLMode:  cfg.DBSSLMode,
	// })
	// if err != nil {
	// 	log.Fatalf("failed opening connection to database: %v", err)
	// }
	// defer sql.Close()

	dbCfg, err := pgxpool.ParseConfig("postgres://user:password@db:5432/db")
	if err != nil {
		log.Fatalf("failed parsing database config: %v", err)
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), dbCfg)
	if err != nil {
		log.Fatalf("failed creating pool: %v", err)
	}
	defer pool.Close()

	e := echo.New()

	e.HTTPErrorHandler = internal_errors.ErrorHandler

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

	oapi.RegisterHandlers(e, internal.Wire(pool))

	e.Logger.Fatal(e.Start(":1323"))
}
