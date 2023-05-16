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

// errorHTML returns HTML page rendered with error message.
func errorHTML(c *gin.Context, code int, message string) {
	c.HTML(code, "error.html", gin.H{
		"code": code,
		"data": message,
	})
}
