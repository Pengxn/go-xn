package controller

import (
	"github.com/gin-gonic/gin"
)

// AddPost will add post.
func AddPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
