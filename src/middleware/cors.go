package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

// CORS creates a new CORS Gin middleware with default options.
// It is a wrapper of cors.Cors handler.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.Default().HandlerFunc(c.Writer, c.Request)
	}
}
