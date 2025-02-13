package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"

	"github.com/Pengxn/go-xn/src/config"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse JWT token from request `Authorization` header
		token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor,
			func(token *jwt.Token) (any, error) {
				return []byte(config.Config.Server.JwtToken), nil
			})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "authorization failed",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "token is invalid",
			})
			return
		}

		c.Set("username", claims["username"])
		c.Set("uid", claims["uid"])
		c.Next()
	}
}
