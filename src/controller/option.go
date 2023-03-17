package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/model"
)

// ListOptions returns all options.
// Request sample:
//
//	GET => /options
func ListOptions(c *gin.Context) {
	options := map[string]string{}
	for _, option := range model.GetAllOptions() {
		options[option.Name] = option.Value
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": options,
	})
}

// GetOption returns an option by 'name' param.
// Request sample:
//
//	GET => /option/:name
func GetOption(c *gin.Context) {
	has, option := model.GetOptionByName(c.Param("name"))

	if has {
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

// InsertOption inserts an option.
// Request sample:
//
//	POST => /option/:name?value=foo
func InsertOption(c *gin.Context) {
	option := &model.Option{
		Name:  c.Query("name"),
		Value: c.Query("value"),
	}

	if model.OptionExist(option.Name) {
		c.JSON(400, gin.H{
			"code":  400,
			"error": "Option you requested to insert already exists.",
		})
	} else {
		if model.AddToOption(option) {
			c.JSON(201, gin.H{
				"code": 201,
				"data": "Insert option data to DB successfully.",
			})
		} else {
			c.JSON(500, gin.H{
				"code":  500,
				"error": "Internal server error occurred when inserting option.",
			})
		}
	}
}

// UpdateOption updates an option.
// Request sample:
//
//	PUT => /option/:name?value=foo1
func UpdateOption(c *gin.Context) {
	option := &model.Option{
		Name:  c.Param("name"),
		Value: c.Query("value"),
	}

	if model.OptionExist(option.Name) {
		c.JSON(400, gin.H{
			"code":  400,
			"error": "Option you requested to update don't exists.",
		})
	} else {
		if model.UpdateOptionByName(option) {
			c.JSON(200, gin.H{
				"code": 200,
				"data": "Update option data successfully.",
			})
		} else {
			c.JSON(500, gin.H{
				"code":  500,
				"error": "Internal server error occurred when updating option.",
			})
		}
	}
}

// DeleteOption deletes option by 'name' param.
// Request sample:
//
//	DELETE => /option/:name
func DeleteOption(c *gin.Context) {
	name := c.Param("name")

	if !model.OptionExist(name) {
		c.JSON(400, gin.H{
			"code":  400,
			"error": "Option you requested to delete don't exists.",
		})
	} else {
		if model.DeleteOptionByName(name) {
			c.JSON(200, gin.H{
				"code": 200,
				"data": "Delete option data successfully.",
			})
		} else {
			c.JSON(500, gin.H{
				"code":  500,
				"error": "Internal server error occurred when deleting option.",
			})
		}
	}
}
