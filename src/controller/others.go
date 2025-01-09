package controller

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
	"golang.org/x/net/idna"

	"github.com/Pengxn/go-xn/src/lib/markdown"
	"github.com/Pengxn/go-xn/src/lib/whois"
	"github.com/Pengxn/go-xn/src/util/log"
)

// GwtWhoisInfo gets domain whois information.
// Request sample:
//
//	GET => /whois?domain=xn--02f.com
func GwtWhoisInfo(c *gin.Context) {
	domain := strings.TrimSpace(c.Query("domain"))
	domain = strings.TrimSuffix(domain, ".") // Trim the dot at the end
	if len(strings.Split(domain, ".")) < 2 { // Need a TLD and a domain body
		c.String(403, "Param (domain="+domain+") is invaild")
		return
	}
	// Convert domain to punycode if it includes non-ASCII characters
	domain, err := idna.ToASCII(domain)
	if err != nil {
		log.Errorf("convert punycode error: %+v, domain: %s", err, domain)
		c.String(403, err.Error())
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
//
//	POST => /upload/upic?file=...
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
	// RSS MIME, refer to https://www.rssboard.org/rss-mime-type-application.txt
	c.Header("Content-Type", "application/rss+xml")
	c.String(200, rss)
}

func Atom(c *gin.Context) {
	rss, err := feed().ToAtom()
	if err != nil {
		log.Errorf("Generate atom content error: %+v", err)
		c.XML(500, "")
		return
	}
	// Atom MIME, refer to https://datatracker.ietf.org/doc/html/rfc4287#section-7
	c.Header("Content-Type", "application/atom+xml")
	c.String(200, rss)
}

func Feed(c *gin.Context) {
	rss, err := feed().ToJSON()
	if err != nil {
		log.Errorf("Generate feed content error: %+v", err)
		c.JSON(500, "")
		return
	}
	// JSON Feed MIME
	// refer to https://jsonfeed.org/version/1.1#suggestions-for-publishers
	c.Header("Content-Type", "application/feed+json")
	c.String(200, rss)
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
		Copyright:   fmt.Sprintf("Copyright (c) 2020-%d The Go-xn Author", now.Year()),
		Image:       &feeds.Image{},
	}
}

// Mdcat renders the markdown page to HTML.
func Mdcat(c *gin.Context) {
	content, err := os.ReadFile("README.md")
	if err != nil {
		log.Errorf("Read README.md error: %+v", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": "Read README.md failed",
		})
		return
	}

	html, err := markdown.ToHTML(content)
	if err != nil {
		log.Errorf("Convert markdown to HTML error: %+v", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": "Convert markdown to HTML failed",
		})
		return
	}

	c.HTML(200, "mdcat.html", gin.H{
		"code": 200,
		"site": map[string]any{
			"title":       "Feng",
			"author":      "Feng.YJ",
			"description": "✍ The platform for publishing and running your blog. [WIP]",
			"html":        template.HTML(html),
		},
	})
}
