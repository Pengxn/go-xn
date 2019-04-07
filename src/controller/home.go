package controller

import (
	"github.com/gin-gonic/gin"

	"go-xn/src/config"
)

// Home and index
func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":     "pong",
		"databaseUrl": config.DBUrl(),
	})
}
