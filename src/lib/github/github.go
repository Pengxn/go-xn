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
	defaultOwner  = "Pengxn"
	defaultRepo   = "go-xn"
	defaultBranch = "main"
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

// GetActionsArtifactLink returns the latest artifact link of the workflow run from GitHub Actions.
// It requires the owner and repo name of the repository, and the branch name.
// The artifact link is the download URL of the artifact file for the current os and arch.
func GetActionsArtifactLink() (string, error) {
	client := github.NewClient(nil)
	ctx := context.Background()

	// API doc url: https://docs.github.com/en/rest/actions/workflow-runs#list-workflow-runs-for-a-repository
	option := &github.ListWorkflowRunsOptions{Branch: defaultBranch}
	runs, _, err := client.Actions.ListRepositoryWorkflowRuns(ctx, defaultOwner, defaultRepo, option)
	if err != nil {
		return "", err
	}

	if len(runs.WorkflowRuns) == 0 {
		return "", errors.New("no latest workflow runs found")
	}

	runID := runs.WorkflowRuns[0].GetID()

	// API doc url: https://docs.github.com/en/rest/actions/artifacts#list-workflow-run-artifacts
	artifacts, _, err := client.Actions.ListWorkflowRunArtifacts(ctx, defaultOwner, defaultRepo, runID, nil)
	if err != nil {
		return "", err
	}

	var archive string
	for _, artifact := range artifacts.Artifacts {
		if artifact.GetName() == runtime.GOOS+"-"+runtime.GOARCH {
			archive = artifact.GetArchiveDownloadURL()
			break
		}
	}

	if archive == "" {
		return "", errors.New("no latest artifact link found")
	}

	return archive, nil
}
