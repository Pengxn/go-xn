package github

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/google/go-github/v69/github"
)

var (
	defaultOwner = "Pengxn"
	defaultRepo  = "go-xn"
)

// Render renders the markdown text to HTML by GitHub API.
// Default mode is "gfm" (GitHub Flavored Markdown),
// refer to https://github.github.com/gfm/ for more details.
func Render(text string) (string, error) {
	client := github.NewClient(nil)

	// API doc url: https://docs.github.com/en/rest/markdown/markdown
	res, _, err := client.Markdown.Render(context.Background(), text, &github.MarkdownOptions{
		Mode:    "gfm",
		Context: defaultOwner + "/" + defaultRepo,
	})
	if err != nil {
		return "", err
	}

	return res, nil
}

// GetLatestAssetLink returns the latest asset link of the release from GitHub.
// It requires the owner and repo name of the repository.
// The asset link is the download URL of the asset file for the current os and arch.
func GetLatestAssetLink() (string, error) {
	client := github.NewClient(nil)

	// API doc url: https://docs.github.com/en/rest/releases/releases#get-the-latest-release
	rel, _, err := client.Repositories.GetLatestRelease(context.Background(), defaultOwner, defaultRepo)
	if err != nil {
		return "", err
	}

	substr := fmt.Sprintf("-%s-%s.", runtime.GOOS, runtime.GOARCH)
	for _, asset := range rel.Assets {
		if strings.Contains(asset.GetName(), substr) {
			return asset.GetBrowserDownloadURL(), nil
		}
	}

	return "", errors.New("no latest asset link found")
}
