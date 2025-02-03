package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/controller"
)

// articlesRoutes registers routes about articles.
func articlesRoutes(g *gin.Engine) {
	g.Any("/articles", controller.ListArticles)

	article := g.Group("/article")
	article.GET("/:slug", controller.GetArticle)
	article.POST("/", controller.InsertArticle)
	article.PUT("/:slug", controller.UpdateArticle)
	article.DELETE("/:slug", controller.DeleteArticle)
}
