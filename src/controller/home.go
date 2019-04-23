package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"go-xn/src/model"
)

// HomePage return home and index page JSON information.
func HomePage(c *gin.Context) {
	articles := model.HomeView()

	c.JSON(200, gin.H{
		"data": articles,
	})
}

// ByPageNum return JSON information by page number
func ByPageNum(c *gin.Context) {
	pageNum := c.Param("pageNum")
	num, _ := strconv.ParseUint(pageNum, 10, 64)

	c.JSON(200, gin.H{
		"pageNum": num,
	})
}
