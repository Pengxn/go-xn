package controller

import (
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/lib/whois"
	"github.com/Pengxn/go-xn/src/util/log"
)

// GwtWhoisInfo gets domain whois information.
// Request sample:
//     GET => /whois?domain=xn--02f.com
func GwtWhoisInfo(c *gin.Context) {
	domain := strings.TrimSpace(c.Query("domain"))
	if strings.HasSuffix(domain, ".") { // Trim the dot at the end
		domain = domain[:len(domain)-1]
	}

	if len(strings.Split(domain, ".")) < 2 { // Need a TLD and a domain body
		c.String(403, "Param (domain="+domain+") is invaild")
		return
	}

	res, err := whois.GetWhois(domain)
	if err != nil {
		log.Errorf("Get Whois Information error: %+v, domain: %s", err, domain)
		c.String(404, err.Error())
		return
	}
	c.String(200, res)
}

// UploadFileForUPic uploads files to the specified file path.
// Request sample:
//     POST => /upload/upic?file=...
func UploadFileForUPic(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Errorf("Get uploaded file error: %+v", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": "Get uploaded file failed",
		})
		return
	}

	if err = c.SaveUploadedFile(file, "./uPic/"+file.Filename); err != nil {
		log.Errorf("Save file uploaded to uPic error: %+v", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": "Save uploaded file failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": map[string]string{
			"url": c.Request.Host + "/upic/" + file.Filename,
		},
	})
}
