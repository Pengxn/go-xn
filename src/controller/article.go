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
	c.JSON(200, gin.H{
		"articles": "Add articles",
	})
}

// UpdateArticle update an articles JSON.
func UpdateArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"articles": "Update a article",
	})
}

// DeleteArticle delete an articles JSON.
func DeleteArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"articles": "Delete articles",
	})
}
