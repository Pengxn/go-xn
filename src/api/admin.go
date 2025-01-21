package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/lib/jwt"
)

type adminAPI struct{}

func NewAdminAPI() *adminAPI {
	return &adminAPI{}
}

func (a *adminAPI) Token(c *gin.Context) {
	claims := jwt.NewClaims(1, "admin")
	token, err := jwt.TokenFromClaims(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "token generation failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"token": token,
	})
}

func (a *adminAPI) RegisterAdmin(c *gin.Context) {
	// TODO: parse request body data

	c.JSON(200, gin.H{})
}
