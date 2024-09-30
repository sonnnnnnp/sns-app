package main

import (
	"github.com/sonnnnnnp/sns-app/pkg/config"
	"github.com/sonnnnnnp/sns-app/pkg/server"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	server.Run(cfg)
}
