package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"

	"github.com/Pengxn/go-xn/src/util/log"
)

// WebdavMethods is a list of supported WebDAV methods.
var WebdavMethods = []string{
	"OPTIONS", "GET", "HEAD", "POST", "DELETE", "PUT",
	"MKCOL", "COPY", "MOVE", "LOCK", "UNLOCK", "PROPFIND", "PROPPATCH",
}

func WebDAV(c *gin.Context) {
	h := webdav.Handler{
		Prefix:     "/dav",
		FileSystem: webdav.Dir("data/webdav"),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			if err != nil {
				log.Errorf("[WebDAV] %s => %s error: %v", r.Method, r.URL.Path, err)
			}
		},
	}
	h.ServeHTTP(c.Writer, c.Request)
}
