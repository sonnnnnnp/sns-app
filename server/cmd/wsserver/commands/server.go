package commands

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/sonnnnnnp/reverie/server/adapter/middleware"
	"github.com/sonnnnnnp/reverie/server/adapter/ws/api"
	"github.com/sonnnnnnp/reverie/server/adapter/ws/controller"
	"github.com/sonnnnnnp/reverie/server/config"
	"github.com/sonnnnnnp/reverie/server/pkg/errors"
	"github.com/sonnnnnnp/reverie/server/pkg/ws"
)

func Run(cfg *config.Config) {
	e := echo.New()

	e.HTTPErrorHandler = errors.ErrorHandler

	middlewarecfg := echomiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet},
	}

	websocket := ws.NewHub()
	go websocket.Listen()

	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.CORSWithConfig(middlewarecfg))
	e.Use(middleware.Config(cfg))
	e.Use(middleware.WebSocket(websocket))

	api.RegisterHandlers(e, controller.New())

	e.Logger.Fatal(e.Start(":4649"))
}
