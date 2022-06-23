package controller

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"

	"github.com/Pengxn/go-xn/src/lib/whois"
	"github.com/Pengxn/go-xn/src/util/log"
)

// GwtWhoisInfo gets domain whois information.
// Request sample:
//     GET => /whois?domain=xn--02f.com
func GwtWhoisInfo(c *gin.Context) {
	domain := strings.TrimSpace(c.Query("domain"))
	domain = strings.TrimSuffix(domain, ".") // Trim the dot at the end
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

	// Save uploaded files to data/uPic directory.
	if err = c.SaveUploadedFile(file, "data/uPic/"+file.Filename); err != nil {
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

func RSS(c *gin.Context) {
	rss, err := feed().ToRss()
	if err != nil {
		log.Errorf("Generate RSS content error: %+v", err)
		c.XML(500, "")
		return
	}
	c.XML(200, rss)
}

func Atom(c *gin.Context) {
	rss, err := feed().ToAtom()
	if err != nil {
		log.Errorf("Generate atom content error: %+v", err)
		c.XML(500, "")
		return
	}
	c.XML(200, rss)
}

func Feed(c *gin.Context) {
	rss, err := feed().ToJSON()
	if err != nil {
		log.Errorf("Generate feed content error: %+v", err)
		c.JSON(500, "")
		return
	}
	c.JSON(200, rss)
}

func feed() *feeds.Feed {
	now := time.Now()
	return &feeds.Feed{
		Title:       "Go-xn",
		Link:        &feeds.Link{},
		Description: "The platform for publishing and running your blog.",
		Author:      &feeds.Author{},
		Updated:     now,
		Created:     now,
		Id:          "",
		Subtitle:    "",
		Items:       []*feeds.Item{},
		Copyright:   "",
		Image:       &feeds.Image{},
	}
}
