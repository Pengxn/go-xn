package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/controller"
)

// othersRoutes registers other routes.
func othersRoutes(g *gin.Engine) {
	// Register routes and methods for uPic, more information
	// to https://blog.svend.cc/upic/tutorials/custom
	// Request sample:
	//     POST => /upload/upic?file=...
	g.POST("/upload/upic", controller.UploadFileForUPic)
	g.Static("/upic", "uPic")
}
