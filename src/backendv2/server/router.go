package server

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"jasoncoding.com/backendv2/config"
	"jasoncoding.com/backendv2/controllers"
	"jasoncoding.com/backendv2/cool"
)

// These routes are only available for the main jasoncoding website
func websiteRoutes(router *gin.Engine) {
	webGroup := router.Group("/")

	// Override and cors in dev
	if config.Cfg.GetString("ENVIRONMENT") == "production" {
		webGroup.Use(cors.New(cors.Options{
			AllowedOrigins: []string{"https://jasoncoding.com"},
		}))
	}

	webGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	webGroup.POST("/contact", controllers.VerifyChallenge(cool.ActGetContact), controllers.GetIdentity)
	webGroup.GET("/challenge/:action", controllers.GetCoolChallenge)
}

// These routes are available for public usage
func publicRoutes(router *gin.Engine) {
	pubGroup := router.Group("/api")
	// ? Place routers down down here if you want public APIs
	pubGroup.GET("/pong", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
}

func RegisterRoutes(router *gin.Engine) {
	router.Use(cors.New(cors.Options{
		AllowCredentials:   true,
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "HEAD", "OPTIONS"},
		OptionsPassthrough: true,
		Debug:              true,
	}))

	websiteRoutes(router)
	publicRoutes(router)
}
