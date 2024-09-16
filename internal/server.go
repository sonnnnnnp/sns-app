package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"

	"github.com/sonnnnnnp/sns-app/internal/adapter/controller"
	"github.com/sonnnnnnp/sns-app/internal/adapter/gateway/db"
	"github.com/sonnnnnnp/sns-app/internal/adapter/middleware"
	"github.com/sonnnnnnp/sns-app/pkg/config"
	"github.com/sonnnnnnp/sns-app/pkg/oapi"
)

type Response struct {
	OK   bool        `json:"ok"`
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type ErrorMessage struct {
	Message interface{} `json:"message"`
}

func Init(cfg *config.Config) {
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

	controller := controller.New(db)

	e := echo.New()

	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			ctx.JSON(http.StatusOK, &Response{
				Code: he.Code,
				OK:   false,
				Data: &ErrorMessage{
					Message: he.Message,
				},
			})
			return
		}
		ctx.JSON(http.StatusOK, &Response{
			Code: http.StatusInternalServerError,
			OK:   false,
			Data: &ErrorMessage{
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

	oapi.RegisterHandlers(e, controller)

	e.Logger.Fatal(e.Start(":1323"))
}
