package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var headerXRequestID = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get id from request
		rid := c.GetHeader(headerXRequestID)
		if rid == "" {
			rid = uuid.New().String()
			c.Request.Header.Add(headerXRequestID, rid)
		}
		// Set the request id to ensure that it is in the response
		c.Writer.Header()[headerXRequestID] = []string{rid}
		c.Next()
	}
}
