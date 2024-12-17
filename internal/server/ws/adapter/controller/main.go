package controller

import (
	"github.com/sonnnnnnp/reverie/internal/server/ws/adapter/api"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

var _ api.ServerInterface = (*Controller)(nil)
