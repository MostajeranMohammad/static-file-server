package main

import (
	"log"

	"github.com/MostajeranMohammad/static-file-server/config"
	"github.com/MostajeranMohammad/static-file-server/internal/application"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	application.Run(*cfg)
}
