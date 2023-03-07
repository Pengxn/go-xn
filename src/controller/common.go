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
