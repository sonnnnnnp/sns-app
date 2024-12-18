package main

import (
	"log"

	"github.com/sonnnnnnp/reverie/server/cmd/httpserver/commands"
	"github.com/sonnnnnnp/reverie/server/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed initializing config: %v", err)
	}

	commands.Run(cfg)
}
