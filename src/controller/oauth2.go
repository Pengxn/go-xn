package controller

import (
	"encoding/json"
	"io"
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

// OAuth2Callback handles the OAuth2 callback from the provider.
func OAuth2Callback(c *gin.Context) {
	code := c.Query("code")
	state := c.Query("state")
	if code == "" || state == "" {
		c.JSON(400, gin.H{
			"message": "missing code or state",
		})
		return
	}

	// Check `state` for CSRF protection
	// TODO: validate state from session or redis
	if state != "todo-random-state" {
		c.JSON(400, gin.H{
			"message": "invalid state",
		})
		return
	}

	oauth2Config := oauth2.Config{
		ClientID:     config.Config.OAuth2.GithubClientID,
		ClientSecret: config.Config.OAuth2.GithubClientSecret,
		Endpoint:     github.Endpoint,
	}

	token, err := oauth2Config.Exchange(c, code)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to exchange token",
		})
		return
	}

	// Use the token to get user info
	// https://docs.github.com/en/rest/users/emails?apiVersion=2022-11-28#list-email-addresses-for-the-authenticated-user
	client := oauth2Config.Client(c, token)
	resp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to get user info",
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		c.JSON(500, gin.H{
			"message": "failed to get user info",
		})
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to read response body",
		})
		return
	}

	// TODO: Handle user info (e.g., create user session)

	c.JSON(200, gin.H{
		"message": "success",
		"data":    json.RawMessage(body),
	})
}
