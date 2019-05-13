package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go-xn/src/model"
)

// ListOptions return all option JSON
func ListOptions(c *gin.Context) {
	options := model.GetAllOptions()

	c.JSON(200, gin.H{
		"code": 200,
		"data": options,
	})
}

// GetOption reutn option information JSON
func GetOption(c *gin.Context) {
	name := c.Param("name")

	has, option := model.GetOptionByName(name)

	if has == true {
		c.JSON(200, gin.H{
			"code": 200,
			"data": option,
		})
	} else {
		c.JSON(404, gin.H{
			"code":  404,
			"error": "The option don't exist.",
		})
	}
}

// InsertOption will add option
func InsertOption(c *gin.Context) {
	option := &model.Option{
		Name:  c.Query("name"),
		Value: c.Query("value"),
	}

	if model.OptionExist(option.Name) == true {
		c.JSON(400, gin.H{
			"code":  400,
			"error": "Option you requested to insert already exists.",
		})
	}

	affected := model.AddToOption(option)

	fmt.Println(affected)
}

// UpdateOption will update option
func UpdateOption(c *gin.Context) {
	//
}

// DeleteOption will delete option
func DeleteOption(c *gin.Context) {
	//
}
