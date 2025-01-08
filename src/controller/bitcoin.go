package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/model"
)

// BitcoinAliases returns the Bitcoin BIP15 alias for specified user.
// Note: The status of standard BIP15 aliases is deferred.
//
// Any => /bitcoin-alias/?handle=fengyj
func BitcoinAliases(c *gin.Context) {
	// TODO: deal with query parameters for different handles.
	has, alias, err := model.GetOptionByName("bitcoin-alias")
	if err != nil {
		errorJSON(c, 500, "failed to get bitcoin alias")
		return
	}
	if !has || alias.Value == "" {
		c.String(404, "the bitcoin alias don't exist")
		return
	}

	c.String(200, alias.Value)
}
