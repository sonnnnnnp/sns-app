package main

import (
	"log"

	"github.com/sonnnnnnp/sns-app/pkg/config"
	"github.com/sonnnnnnp/sns-app/pkg/server"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed initializing config: %v", err)
	}

	server.Run(cfg)
}
