package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"jasoncoding.com/backendv2/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/getcontact", controllers.GetIdentity)

	return router
}
