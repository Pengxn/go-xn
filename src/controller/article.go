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

// GetArticle list all articles.
func GetArticle(c *gin.Context) {
	id := c.Param("id")

	c.JSON(200, gin.H{
		"code": 200,
		"data": id,
	})
}

// AddArticle list all articles.
func AddArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"articles": "Add articles",
	})
}

// UpdateArticle list all articles.
func UpdateArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"articles": "Update a article",
	})
}

// DeleteArticle list all articles.
func DeleteArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"articles": "Delete articles",
	})
}
