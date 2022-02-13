package server

import (
	"log"

	"jasoncoding.com/backendv2/config"
)

func Init() {
	cfg := config.GetConfig()
	router := NewRouter()

	port := cfg.GetString("PORT")
	log.Printf("Server will run on port %s\n", port)
	router.Run(":" + port)
}
