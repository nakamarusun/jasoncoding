package server

import (
	"fmt"
	"log"

	"jasoncoding.com/backendv2/config"
)

func Init() {
	cfg := config.GetConfig()
	router := NewRouter()

	// Gets the port based on configuration
	port := cfg.GetString("PORT")
	log.Printf("Server will run on port %s. Environment is '%s'\n", port, cfg.GetString("ENVIRONMENT"))

	// Gets the address based on environment
	addr := ""
	if cfg.GetString("ENVIRONMENT") == "development" {
		addr = "127.0.0.1"
	}

	router.Run(fmt.Sprintf("%s:%s", addr, port))
}
