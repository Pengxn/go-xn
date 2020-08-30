package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/model"
)

// ListArticles returns the number of all articles.
// Request sample:
//     GET => /articles
func ListArticles(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":  200,
		"count": model.ArticlesCount(),
	})
}

// GetArticle gets an article by 'id' param.
// Request sample:
//     GET => /article/1
func GetArticle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	article, _ := model.ArticleByID(id)

	c.JSON(200, gin.H{
		"code":    200,
		"article": article,
	})
}

// InsertArticle inserts an article.
// Request sample:
//     POST => /article?title=foo&status=1&content=bar
func InsertArticle(c *gin.Context) {
	status, _ := strconv.Atoi(c.DefaultQuery("status", "0"))
	article := &model.Article{
		Title:   c.Query("title"),
		Status:  status,
		Content: c.Query("content"),
	}

	if model.AddArticle(article) {
		c.JSON(201, gin.H{
			"code": 201,
			"data": "Insert article data successfully.",
		})
	} else {
		c.JSON(500, gin.H{
			"code":  500,
			"error": "Internal server error occurred when inserting article.",
		})
	}
}

// UpdateArticle updates an article.
func UpdateArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "Update an Article",
		"data": c.Param("id"),
	})
}

// DeleteArticle deletes an article.
func DeleteArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "Delete an Article",
		"data": c.Param("id"),
	})
}
