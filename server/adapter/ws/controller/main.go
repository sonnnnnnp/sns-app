package controller

import (
	"github.com/sonnnnnnp/reverie/server/adapter/ws/api"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

var _ api.ServerInterface = (*Controller)(nil)
