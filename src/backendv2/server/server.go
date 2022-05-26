package server

import (
	"fmt"
	"log"

	"jasoncoding.com/backendv2/config"
	"github.com/gin-gonic/gin"
)

func Init() {

	// Gets the port based on configuration
	port := config.Cfg.GetString("PORT")
	log.Printf("Server will run on port %s. Environment is '%s'\n", port, config.Cfg.GetString("ENVIRONMENT"))

	// Gets the address based on environment
	addr := ""
	if config.Cfg.GetString("ENVIRONMENT") == "development" {
		addr = "127.0.0.1"
	}

	// Create new gin server
	router := gin.New()
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	RegisterRoutes(router)

	router.Run(fmt.Sprintf("%s:%s", addr, port))
}
