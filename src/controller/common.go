package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Redirect returns a HTTP redirect to the specific route path.
func Redirect(routePath string) func(*gin.Context) {
	return func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, routePath)
	}
}

const (
	errorCode   = 0   // common error
	successCode = 200 // success code
)

// errorHTML returns HTML page rendered with error message.
func errorHTML(c *gin.Context, status int, message string) {
	c.HTML(status, "error.html", gin.H{
		"code":    status,
		"message": message,
	})
}

// errorJSON returns JSON response with error message.
func errorJSON(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"code":    errorCode,
		"message": message,
	})
}

// dataJSON returns JSON response with data.
func dataJSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{
		"code": successCode,
		"data": data,
	})
}
