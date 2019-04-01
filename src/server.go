package app

import (
	"github.com/gin-gonic/gin"

	"go-xn/src/route"
)

// Run is the entry point to the server app.
// Parses the arguments slice and routes.
func Run() {
	g := gin.Default()

	route.InitRoutes(g)
	route.InitStatic(g)

	g.Run(":3000")
}
