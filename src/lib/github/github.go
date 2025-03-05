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

// GetNightlyLink returns the nightly build artifact link from nightly.link.
func GetNightlyLink() (string, error) {
	_, artifactName, err := GetActionsArtifactLink()
	if err != nil {
		return "", err
	}

	// nightly.link is a service to provide nightly build artifact download link.
	nightlyURL := "https://nightly.link/Pengxn/go-xn/workflows/test/main"

	return fmt.Sprintf("%s/%s.zip", nightlyURL, artifactName), nil
}

// GetActionsArtifactLink returns the latest artifact link of the workflow run from GitHub Actions.
// It requires the owner and repo name of the repository, and the branch name.
// The artifact link is the download URL of the artifact file for the current os and arch.
//
// But it requires authentication with `actions:read` scope to access the archived artifacts links.
func GetActionsArtifactLink() (string, string, error) {
	client := github.NewClient(nil)
	ctx := context.Background()

	// API doc url: https://docs.github.com/en/rest/actions/workflow-runs#list-workflow-runs-for-a-repository
	option := &github.ListWorkflowRunsOptions{Branch: defaultBranch}
	runs, _, err := client.Actions.ListRepositoryWorkflowRuns(ctx, defaultOwner, defaultRepo, option)
	if err != nil {
		return "", "", err
	}

	if len(runs.WorkflowRuns) == 0 {
		return "", "", errors.New("no latest workflow runs found")
	}

	runID := runs.WorkflowRuns[0].GetID()

	// TODO: check the status of the workflow run, if it's `in_progress` to detect the latest artifact.
	// If `in_progress` workflow run not found, then get the latest completed workflow run.

	// API doc url: https://docs.github.com/en/rest/actions/artifacts#list-workflow-run-artifacts
	artifacts, _, err := client.Actions.ListWorkflowRunArtifacts(ctx, defaultOwner, defaultRepo, runID, nil)
	if err != nil {
		return "", "", err
	}

	var archive, artifactName string
	for _, artifact := range artifacts.Artifacts {
		if artifact.GetName() == runtime.GOOS+"-"+runtime.GOARCH {
			archive = artifact.GetArchiveDownloadURL()
			artifactName = artifact.GetName()
			break
		}
	}

	if archive == "" {
		return "", "", errors.New("no latest artifact link found")
	}

	return archive, artifactName, nil
}
