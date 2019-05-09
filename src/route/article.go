package route

import (
	"github.com/gin-gonic/gin"

	"go-xn/src/controller"
)

// articlesRoutes register routes about article.
func articlesRoutes(g *gin.Engine) {
	g.Any("/articles", controller.ListArticles)

	article := g.Group("/article")
	article.GET("/:id", controller.GetArticle)
	article.POST("/", controller.InsertArticle)
	article.PUT("/:id", controller.UpdateArticle)
	article.DELETE("/:id", controller.DeleteArticle)
}
