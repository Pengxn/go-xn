package controller

import (
	"github.com/gin-gonic/gin"
)

// Home and index.
func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
