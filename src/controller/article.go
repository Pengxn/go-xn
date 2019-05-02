package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"go-xn/src/model"
)

// ListArticles list all articles JSON.
func ListArticles(c *gin.Context) {
	count := model.ArticlesCount()

	c.JSON(200, gin.H{
		"code":  200,
		"count": count,
	})
}

// GetArticle get an articles JSON.
func GetArticle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	article, _ := model.ArticleByID(id)

	c.JSON(200, gin.H{
		"code":    200,
		"article": article,
	})
}

// AddArticle add an articles JSON.
func AddArticle(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"code": "Add an Article",
		"data": id,
	})
}

// UpdateArticle update an articles JSON.
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"code": "Update an Article",
		"data": id,
	})
}

// DeleteArticle delete an articles JSON.
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"code": "Delete an Article",
		"data": id,
	})
}
