package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/util/log"
)

// uPicRoute register routes and methods foruPic, more information
// to https://blog.svend.cc/upic/tutorials/custom
func uPicRoute(g *gin.Engine) {
	// Request sample:
	//     POST => /upload/upic?file=...
	g.POST("/upload/upic", uploadFileForUPic)
}

// uploadFileForUPic uploads files to the specified file path.
func uploadFileForUPic(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Errorf("Get uploaded file error: %+v", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": "Get uploaded file failed",
		})
	}

	if err = c.SaveUploadedFile(file, "./uPic/"+file.Filename); err != nil {
		log.Errorf("Save file uploaded to uPic error: %+v", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": "Save uploaded file failed",
		})
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": map[string]string{
			"url": c.Request.Host + "/uPic/" + file.Filename,
		},
	})
}
