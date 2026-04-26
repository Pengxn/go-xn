package controller

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"
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
				slog.Error("[WebDAV]", slog.String("method", r.Method), slog.String("path", r.URL.Path), slog.Any("error", err))
			}
		},
	}
	h.ServeHTTP(c.Writer, c.Request)
}
