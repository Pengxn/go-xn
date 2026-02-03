package github

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"strings"

	"github.com/google/go-github/v82/github"
)

var (
	defaultOwner  = "Pengxn"
	defaultRepo   = "go-xn"
	defaultBranch = "main"
)

// Render renders the markdown text to HTML by GitHub API.
// Default mode is "gfm" (GitHub Flavored Markdown),
// refer to [GitHub Flavored Markdown Spec] for more details.
//
// [GitHub Flavored Markdown Spec]: https://github.github.com/gfm/#gfm-overview
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
	// release asset name suffix, e.g. <name>-<version>-linux-amd64.tar.gz
	substr := fmt.Sprintf("-%s-%s.", runtime.GOOS, runtime.GOARCH)
	for _, asset := range rel.Assets {
		if strings.Contains(asset.GetName(), substr) {
			return asset.GetBrowserDownloadURL(), nil
		}
	}

	return "", errors.New("no latest asset link found")
}

// GetNightlyLink returns the nightly build artifact link from [nightly.link].
//
// [nightly.link]: https://nightly.link
func GetNightlyLink() (string, error) {
	_, artifactName, err := GetActionsArtifactLink()
	if err != nil {
		return "", err
	}

	// nightly.link is a service to provide nightly build artifact download link.
	nightlyURL := "https://nightly.link/Pengxn/go-xn/workflows/test"

	return fmt.Sprintf("%s/%s/%s.zip", nightlyURL, defaultBranch, artifactName), nil
}

// GetActionsArtifactLink returns the latest artifact link of the workflow run from GitHub Actions.
// It requires the owner and repo name of the repository, and the branch name.
// The artifact link is the download URL of the artifact file for the current os and arch.
//
// But it requires authentication with `actions:read` scope to access the archived artifacts links.
func GetActionsArtifactLink() (string, string, error) {
	client := github.NewClient(nil)
	ctx := context.Background()

	// Get the latest workflow run from list of workflow runs.
	// API doc url: https://docs.github.com/en/rest/actions/workflow-runs#list-workflow-runs-for-a-repository
	option := &github.ListWorkflowRunsOptions{Branch: defaultBranch}
	runs, _, err := client.Actions.ListRepositoryWorkflowRuns(ctx, defaultOwner, defaultRepo, option)
	if err != nil {
		return "", "", err
	}

	var runID int64
	for _, run := range runs.WorkflowRuns {
		// Skip the `Dependabot Updates` workflow runs, as they are not the nightly build.
		// Its complete path is `dynamic/dependabot/dependabot-updates`.
		if strings.Contains(run.GetPath(), "dependabot") {
			continue
		}
		if run.GetStatus() == "completed" {
			runID = run.GetID()
			break
		}
	}

	if runID == 0 {
		return "", "", errors.New("no latest completed workflow runs found")
	}

	// Get the list of artifacts for specific workflow run.
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
