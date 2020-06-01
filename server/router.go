package server

import "github.com/gin-gonic/gin"

func setupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/", landingPage)

	// Endpoints for Shopping your dress
	shop := v1.Group("/shop")
	shop.GET("/home", shopHome)

}
