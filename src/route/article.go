package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/controller"
)

// articlesRoutes registers routes about articles.
func articlesRoutes(g *gin.Engine) {
	g.Any("/articles", controller.ListArticles)

	article := g.Group("/article")
	article.GET("/:url", controller.GetArticle)
	article.POST("/", controller.InsertArticle)
	article.PUT("/:id", controller.UpdateArticle)
	article.DELETE("/:id", controller.DeleteArticle)
}
