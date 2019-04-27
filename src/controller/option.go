package controller

import (
	"github.com/gin-gonic/gin"

	"go-xn/src/model"
)

// ListOptions return all option JSON
func ListOptions(c *gin.Context) {
	//
}

// GetOption reutn option information JSON
func GetOption(c *gin.Context) {
	name := c.Param("name")

	option := model.OptionByName(name)

	c.JSON(200, gin.H{
		"code":   200,
		"option": option,
	})
}

// AddOption will add option
func AddOption(c *gin.Context) {
	//
}

// UpdateOption will update option
func UpdateOption(c *gin.Context) {
	//
}

// DeleteOption will delete option
func DeleteOption(c *gin.Context) {
	//
}
