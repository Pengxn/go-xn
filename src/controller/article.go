package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/model"
)

// ListArticles return the number of all articles
// Request sample:
//     GET /articles
func ListArticles(c *gin.Context) {
	count := model.ArticlesCount()

	c.JSON(200, gin.H{
		"code":  200,
		"count": count,
	})
}

// GetArticle get an articles by 'id' param
// Request sample:
//     GET /article/1
func GetArticle(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	article, _ := model.ArticleByID(id)

	c.JSON(200, gin.H{
		"code":    200,
		"article": article,
	})
}

// InsertArticle will insert an articles
// Request sample:
//     POST /article?title=foo&status=1&content=bar
func InsertArticle(c *gin.Context) {
	title := c.Query("title")
	statusString := c.DefaultQuery("status", "0")
	content := c.Query("content")

	status, _ := strconv.Atoi(statusString)

	article := &model.Article{
		Title:   title,
		Status:  status,
		Content: content,
	}

	if model.AddArticle(article) == true {
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
