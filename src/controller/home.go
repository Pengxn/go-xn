package controller

import (
	"github.com/gin-gonic/gin"

	"go-xn/src/model"
)

// Home and index
func Home(c *gin.Context) {
	articles, _ := model.HomeView()

	c.JSON(200, gin.H{
		"code":     200,
		"articles": articles,
	})
}
