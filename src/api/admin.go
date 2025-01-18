package api

import "github.com/gin-gonic/gin"

type adminAPI struct{}

func NewAdminAPI() *adminAPI {
	return &adminAPI{}
}

func (a *adminAPI) RegisterAdmin(c *gin.Context) {
	// TODO: parse request body data

	c.JSON(200, gin.H{})
}
