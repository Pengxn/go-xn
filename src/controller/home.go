package controller

import (
	"github.com/gin-gonic/gin"

	"go-xn/src/model"
)

// Home and index
func Home(c *gin.Context) {
	articles, _ := model.HomeView()

	c.HTML(200, "index.hbs", gin.H{
		"article": articles,
	})
}
