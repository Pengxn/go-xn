package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"go-xn/src/model"
)

// DefaultLimit is default limit number per page
// And other constant
const (
	DefaultLimit int = 8
)

// HomePage return home and index page JSON information.
func HomePage(c *gin.Context) {
	articles := model.HomeView()

	c.JSON(200, gin.H{
		"data": articles,
	})
}

// ArticlesPage return JSON information by page number
func ArticlesPage(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Param("pageNum"))
	limitNum, _ := strconv.Atoi(c.Param("limit"))

	if pageNum == 0 {
		pageNum = 1
	}
	if limitNum == 0 {
		limitNum = DefaultLimit
	}

	articles := model.ArticlesByPage(limitNum, pageNum)

	c.JSON(200, gin.H{
		"code": 200,
		"data": articles,
	})
}
