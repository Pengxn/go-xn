package route

import (
	"github.com/gin-gonic/gin"

	"go-xn/src/controller"
)

// ArticlesRoutes registere routes about article.
func ArticlesRoutes(g *gin.Engine) {
	g.Any("/articles", controller.ListArticles)

	article := g.Group("/article")
	article.GET("/:id", controller.GetArticle)
	article.POST("/:id", controller.AddArticle)
	article.PUT("/:id", controller.UpdateArticle)
	article.DELETE("/:id", controller.DeleteArticle)
}
