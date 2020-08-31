package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/model"
)

// DefaultPageLimit is default limit number per page,
// and other constant.
const (
	DefaultPageLimit int = 8
)

// HomePage returns home and index page.
// Request sample:
//     GET => /
//     GET => /page/1
func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": model.ArticlesByPage(DefaultPageLimit, 1),
	})
}

// ArticlesPage returns JSON information by page number.
// Request sample:
//     GET => /page/1?limit=6
func ArticlesPage(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Param("pageNum"))
	limitNum, _ := strconv.Atoi(c.Query("limit"))

	if pageNum <= 0 {
		pageNum = 1
	}
	if limitNum == 0 {
		limitNum = DefaultPageLimit
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": model.ArticlesByPage(limitNum, pageNum),
	})
}
