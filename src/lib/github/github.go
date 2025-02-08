package github

import (
	"context"

	"github.com/google/go-github/v69/github"
)

// Render renders the markdown text to HTML by GitHub API.
// Default mode is "gfm" (GitHub Flavored Markdown),
// refer to https://github.github.com/gfm/ for more details.
func Render(text string) (string, error) {
	client := github.NewClient(nil)

	// API doc url: https://docs.github.com/en/rest/markdown/markdown
	res, _, err := client.Markdown.Render(context.Background(), text, &github.MarkdownOptions{
		Mode:    "gfm",
		Context: "Pengxn/go-xn",
	})
	if err != nil {
		return "", err
	}

	return res, nil
}
