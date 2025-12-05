package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"

	"github.com/Pengxn/go-xn/src/config"
)

// OAuth2Redirect handles OAuth2 redirection to the provider's authorization page.
func OAuth2Redirect(c *gin.Context) {
	provider := c.Param("provider")
	if provider == "" {
		c.JSON(400, gin.H{
			"message": "unspecified provider",
		})
		return
	}

	// Currently only support GitHub OAuth2
	// refer to: https://docs.github.com/v3/oauth/
	if provider != "github" {
		c.JSON(400, gin.H{
			"message": "unsupported provider",
		})
		return
	}

	if config.Config.OAuth2.GithubClientID == "" {
		c.JSON(500, gin.H{
			"message": "github oauth2 not enabled",
		})
		return
	}

	// Generate state for CSRF protection (in production, store this in session/redis)
	// TODO: Use a secure random state and store it to validate in callback handler
	state := "todo-random-state"

	scheme := "http://"
	if c.Request.Header.Get("X-Forwarded-Proto") == "https" || strings.HasPrefix(c.Request.Referer(), "https://") {
		scheme = "https://"
	}
	redirectURL := scheme + c.Request.Host + "/oauth/callback"

	oauth2Config := oauth2.Config{
		ClientID:     config.Config.OAuth2.GithubClientID,
		ClientSecret: config.Config.OAuth2.GithubClientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     github.Endpoint,

		// Specify minimal scopes to get user email and profile info, refer to:
		// https://docs.github.com/en/apps/oauth-apps/building-oauth-apps/scopes-for-oauth-apps#available-scopes
		Scopes: []string{"user:email", "read:user"},
	}

	// Redirect user to GitHub authorization page
	c.Redirect(http.StatusFound, oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOnline))
}
