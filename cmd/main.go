package main

import (
	server "github.com/sonnnnnnp/sns-app/internal"

	"github.com/sonnnnnnp/sns-app/pkg/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	server.Init(cfg)
}
