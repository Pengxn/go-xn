package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/model"
)

// BasicAuth is a middleware to check basic authentication.
// The auth middleware is used for WebDAV service, refer to:
// http://www.webdav.org/specs/rfc2617.html#n-basic-authentication-scheme
func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok || username == "" || password == "" {
			// Set WWW-Authenticate header to tell the client to authenticate.
			// refer to: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/WWW-Authenticate
			c.Writer.Header().Set("WWW-Authenticate", `Basic realm="fyj WebDAV"`)
			c.String(http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		has, user, err := model.GetUserByName(username)
		if err != nil || !has {
			c.String(http.StatusInternalServerError, "server error")
			c.Abort()
			return
		}

		if username != user.Name || password != user.Password {
			c.String(http.StatusForbidden, "not permitted")
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Next()
	}
}
