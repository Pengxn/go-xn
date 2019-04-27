package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"go-xn/src/model"
)

// DefaultLimit is default limit number per page
// And other constant
const (
	DefaultLimit uint64 = 8
)

// HomePage return home and index page JSON information.
func HomePage(c *gin.Context) {
	articles := model.HomeView()

	c.JSON(200, gin.H{
		"data": articles,
	})
}

// AboutPage return 'about' JSON
func AboutPage(c *gin.Context) {
	option := model.OptionByName("")

	c.JSON(200, gin.H{
		"code":          200,
		"option_name":   option.Name,
		"optioin_value": option.Value,
	})
}

// ByPageNum return JSON information by page number
func ByPageNum(c *gin.Context) {
	pageNum, _ := strconv.ParseUint(c.Param("pageNum"), 10, 64)
	limitNum, _ := strconv.ParseUint(c.Query("limit"), 10, 64)

	if pageNum == 0 {
		pageNum = 1
	}
	if limitNum == 0 {
		limitNum = DefaultLimit
	}

	c.JSON(200, gin.H{
		"pageNum":  pageNum,
		"limitNum": limitNum,
	})
}
