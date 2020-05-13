package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func landingPage(c *gin.Context) {
	c.String(http.StatusOK, "Hello, I am chaitu. The API Developer")
}
