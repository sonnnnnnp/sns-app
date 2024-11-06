package main

import (
	"log"

	"github.com/sonnnnnnp/sns-app/pkg/config"
)

func main() {
	_, err := config.New()
	if err != nil {
		log.Fatalf("failed initializing config: %v", err)
	}
}
