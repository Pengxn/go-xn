package controller

import (
	"github.com/gin-gonic/gin"
)

// ListArticles list all articles JSON.
func ListArticles(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
	})
}

// GetArticle get an articles JSON.
func GetArticle(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"code": 200,
		"data": id,
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
