package controller

import (
	"fmt"
	"net/http"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/gin-gonic/gin"
)

// WebAuthnWellKnown returns well-known json for webauthn.
// The `origins` field is relate with the WebAuthn config.
func WebAuthnWellKnown(c *gin.Context) {
	var origins []string
	for _, origin := range config.Config.WebAuthn.RPOrigins {
		origins = append(origins, fmt.Sprintf("https://%s", origin))
	}

	c.JSON(http.StatusOK, gin.H{
		"origins": origins,
	})
}
