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

// InsertArticle add an articles JSON.
func InsertArticle(c *gin.Context) {
	title := c.Query("title")
	status := c.Query("status")
	content := c.Query("content")

	c.JSON(200, gin.H{
		"code":    200,
		"title":   title,
		"status":  status,
		"content": content,
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
