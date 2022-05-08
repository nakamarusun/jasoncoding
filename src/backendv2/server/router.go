package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"jasoncoding.com/backendv2/controllers"
	"jasoncoding.com/backendv2/config"
)

// These routes are only available for the main jasoncoding website
func websiteRoutes(router *gin.Engine) {
	webGroup := router.Group("/")

	// Override and cors in dev
	if config.GetConfig().GetString("ENVIRONMENT") == "production" {
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowOrigins = []string{"https://jasoncoding.com"}
		webGroup.Use(cors.New(corsConfig))
	} else {
		webGroup.Use(cors.Default())
	}

	webGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	webGroup.POST("/getcontact", controllers.GetIdentity)
}

// These routes are available for public usage
func publicRoutes(router *gin.Engine) {
	pubGroup := router.Group("/api")
	// https://stackoverflow.com/a/56348408/12709867
	pubGroup.Use(cors.Default())
	// ? Place routers down down here if you want public APIs

}

func RegisterRoutes(router *gin.Engine) {
	websiteRoutes(router);
	publicRoutes(router);
}
