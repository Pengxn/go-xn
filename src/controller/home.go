package controller

import (
	"github.com/gin-gonic/gin"

	"go-xn/src/model"
)

// HomePage return home and index page JSON information.
func HomePage(c *gin.Context) {
	articles, _ := model.HomeView()

	c.JSON(200, gin.H{
		"data": articles,
	})
}
