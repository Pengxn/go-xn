package controller

import (
	"html/template"
	"log/slog"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/lib/markdown"
	"github.com/Pengxn/go-xn/src/model"
)

// DefaultPageLimit is default limit number per page,
// and other constant.
const (
	DefaultPageLimit int = 8
)

// HomePage returns home and index page.
// Request sample:
//
//	GET => /
//	GET => /page/1
func HomePage(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"code":     200,
		"articles": model.ArticlesByPage(DefaultPageLimit, 1),
		"now":      time.Now(),
		"site": map[string]string{
			"title":       "Feng",
			"author":      "Feng.YJ",
			"description": "✍ The platform for publishing and running your blog. [WIP]",
		},
	})
}

// ArticlesPage returns JSON information by page number.
// Request sample:
//
//	GET => /page/1?limit=6
func ArticlesPage(c *gin.Context) {
	pageNum, _ := strconv.Atoi(c.Param("pageNum"))
	limitNum, _ := strconv.Atoi(c.Query("limit"))

	if pageNum <= 0 {
		pageNum = 1
	}
	if limitNum == 0 {
		limitNum = DefaultPageLimit
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": model.ArticlesByPage(limitNum, pageNum),
	})
}

// ListArticles returns the number of all articles.
// Request sample:
//
//	GET => /articles
func ListArticles(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":  200,
		"count": model.ArticlesCount(),
	})
}

// GetArticle gets an article by 'slug' param.
// Request sample:
//
//	GET => /article/article-custom-slug-path
func GetArticle(c *gin.Context) {
	article, exist := model.ArticleBySlug(c.Param("slug"))
	if !exist {
		errorHTML(c, 404, "Article Not Found")
		return
	}

	content, err := markdown.ToHTML([]byte(article.Content))
	if err != nil {
		slog.Error("convert markdown to HTML error", slog.Any("error", err))
		c.JSON(500, gin.H{
			"code": 500,
			"data": "Convert markdown to HTML failed",
		})
		return
	}

	c.HTML(200, "mdcat.html", gin.H{
		"code": 200,
		"site": map[string]any{
			"title":       "Feng",
			"author":      "Feng.YJ",
			"description": "✍ The platform for publishing and running your blog. [WIP]",
			"html":        template.HTML(content),
		},
	})
}

// InsertArticle inserts an article.
// Request sample:
//
//	POST => /article?title=foo&status=1&content=bar
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
		return
	}

	c.JSON(500, gin.H{
		"code":  500,
		"error": "Internal server error occurred when inserting article.",
	})
}

// UpdateArticle updates an article.
func UpdateArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "Update an Article",
		"data": c.Param("slug"),
	})
}

// DeleteArticle deletes an article.
func DeleteArticle(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "Delete an Article",
		"data": c.Param("slug"),
	})
}
