package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// MetaWeblog is the controller of MetaWeblog API.
// The MetaWeblog API (MWA) is a programming interface that allows external
// programs to get and set the text and attributes of weblog posts.
//
// More info to http://xmlrpc.com/storyreader/story2198.md and
// https://codex.wordpress.org/XML-RPC_MetaWeblog_API
func MetaWeblog(c *gin.Context) {
	c.Header("Content-Type", "text/xml")

	// TODO: implement MetaWeblog API, return fault response temporarily.
	c.String(200, fmt.Sprintf(faultResponse, 200, "not implemented"))
}

var faultResponse = `<?xml version="1.0" encoding="UTF-8"?>
<methodResponse>
  <fault>
    <value>
      <struct>
        <member>
          <name>faultCode</name>
          <value><int>%d</int></value>
        </member>
        <member>
          <name>faultString</name>
          <value><string>%s</string></value>
        </member>
      </struct>
    </value>
  </fault>
</methodResponse>`
