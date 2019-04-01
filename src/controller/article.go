package controller

import (
	"github.com/gin-gonic/gin"
)

// ListArticles list all artivles.
func ListArticles(c *gin.Context) {
	c.JSON(200, gin.H{
		"articles": "list",
	})
}

// GetArticle list all artivles.
func GetArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"articles": "get a article",
	})
}

// AddArticle list all artivles.
func AddArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"articles": "Add articles",
	})
}

// UpdateArticle list all artivles.
func UpdateArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"articles": "Update a article",
	})
}

// DeleteArticle list all artivles.
func DeleteArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"articles": "Delete articles",
	})
}
