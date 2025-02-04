package middleware

import "github.com/gin-gonic/gin"

func NoIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set the X-Robots-Tag header to noindex,
		// refer to https://developers.google.com/search/reference/robots_meta_tag
		c.Header("X-Robots-Tag", "noindex")
		c.Next()
	}
}
