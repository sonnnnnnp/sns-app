package main

import (
	"github.com/sonnnnnnp/sns-app/internal/app"
	"github.com/sonnnnnnp/sns-app/pkg/config"
)

func main() {
	cfg, err := config.New()
	if err!= nil {
        panic(err)
    }

	app.Init(cfg)
}
