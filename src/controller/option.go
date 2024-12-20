package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"xorm.io/xorm"

	"github.com/Pengxn/go-xn/src/model"
	"github.com/Pengxn/go-xn/src/util/log"
)

// ListOptions returns all options.
// Request sample:
//
//	GET => /options
func ListOptions(c *gin.Context) {
	list, err := model.GetAllOptions()
	if err != nil {
		log.Errorf("ListOptions: %s", err)
		errorJSON(c, 500, "failed to get options")
		return
	}

	options := map[string]string{}
	for _, option := range list {
		options[option.Name] = option.Value
	}

	dataJSON(c, 200, options)
}

// GetOption returns an option by 'name' param.
// Request sample:
//
//	GET => /option/:name
func GetOption(c *gin.Context) {
	name := strings.TrimSpace(c.Param("name"))
	if name == "" {
		errorHTML(c, 400, "option name is required")
		return
	}

	has, option, err := model.GetOptionByName(name)
	if err != nil {
		log.Errorf("GetOption: %s", err)
		errorJSON(c, 500, "failed to get option data")
		return
	}
	if has {
		dataJSON(c, 200, option.Value)
		return
	}
	errorHTML(c, 404, "the option don't exist")
}

// InsertOption inserts an option.
// Request sample:
//
//	POST => /option
func InsertOption(c *gin.Context) {
	name := strings.TrimSpace(c.PostForm("name"))
	value := strings.TrimSpace(c.PostForm("value"))
	if name == "" || value == "" {
		errorJSON(c, 400, "option name and value are required")
		return
	}

	exist, err := model.OptionExist(name)
	if err != nil || exist {
		errorJSON(c, 400, "option you requested to insert already exists")
		return
	}

	success, err := model.AddOption(&model.Option{Name: name, Value: string(value)})
	if err != nil || !success {
		log.Errorf("InsertOption: %s", err)
		errorJSON(c, 500, "failed to insert option data")
		return
	}

	dataJSON(c, 201, "insert option data successfully") // 201 Created
}

// UpdateOption updates an option.
// Request sample:
//
//	PUT => /option/:name
func UpdateOption(c *gin.Context) {
	name := strings.TrimSpace(c.Param("name"))
	value, err := c.GetRawData()
	if err != nil || name == "" {
		errorJSON(c, 400, "option value is required")
		return
	}

	success, err := model.UpdateOptionByName(model.Option{Name: name, Value: string(value)})
	if err == xorm.ErrNotExist {
		errorJSON(c, 400, "option you requested to update don't exists")
		return
	}
	if err != nil || !success {
		log.Errorf("UpdateOption: %s", err)
		errorJSON(c, 500, "update option data failed")
		return
	}

	dataJSON(c, 200, "update option data successfully")
}

// DeleteOption deletes option by 'name' param.
// Request sample:
//
//	DELETE => /option/:name
func DeleteOption(c *gin.Context) {
	name := strings.TrimSpace(c.Param("name"))
	if name == "" {
		errorJSON(c, 400, "option name is required")
		return
	}

	success, err := model.DeleteOptionByName(name)
	if err == xorm.ErrNotExist {
		dataJSON(c, 204, "the option you requested to delete doesn't exist")
		return
	}
	if err != nil || !success {
		log.Errorf("DeleteOption: %s", err)
		errorJSON(c, 500, "delete option data failed")
		return
	}

	dataJSON(c, 200, "delete option data successfully")
}
