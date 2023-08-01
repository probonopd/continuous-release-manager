package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Error: GITHUB_TOKEN environment variable not set.")
		os.Exit(1)
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repoOwner := ""
	repoName := ""
	releaseCommitHash := ""
	releaseName := "continuous"
	releaseTag := "continuous"

	isGitHubActions := os.Getenv("GITHUB_ACTIONS") == "true"
	isCirrusCI := os.Getenv("CIRRUS_CI") == "true"

	if isGitHubActions {
		repoOwner = os.Getenv("GITHUB_REPOSITORY_OWNER")
		repoName = extractRepositoryName(os.Getenv("GITHUB_REPOSITORY"))
		releaseCommitHash = os.Getenv("GITHUB_SHA")
	} else if isCirrusCI {
		repoOwner = os.Getenv("CIRRUS_REPO_OWNER")
		repoName = os.Getenv("CIRRUS_REPO_NAME")
		releaseCommitHash = os.Getenv("CIRRUS_CHANGE_IN_REPO")
	} else {
		fmt.Println("Error: Unsupported CI environment.")
		os.Exit(1)
	}

	logInfo("Starting release management...")
	logVerbose(fmt.Sprintf("Repository Owner: %s", repoOwner))
	logVerbose(fmt.Sprintf("Repository Name: %s", repoName))
	logVerbose(fmt.Sprintf("Release Tag: %s", releaseTag))
	logVerbose(fmt.Sprintf("Release Commit Hash: %s", releaseCommitHash))

	// Check if the release with the name "continuous" already exists
	logInfo("Checking for existing release...")
	release, _, err := client.Repositories.GetReleaseByTag(ctx, repoOwner, repoName, releaseTag)
	if err != nil {
		// An error occurred while retrieving the release
		if _, ok := err.(*github.ErrorResponse); ok && err.(*github.ErrorResponse).Response.StatusCode == 404 {
			// The release does not exist yet, proceed to create it
			logInfo("Release with the name 'continuous' does not exist. Creating a new release...")
			newRelease := &github.RepositoryRelease{
				TagName:         &releaseTag,
				TargetCommitish: &releaseCommitHash,
				Name:            &releaseName,
			}
			createdRelease, _, err := client.Repositories.CreateRelease(ctx, repoOwner, repoName, newRelease)
			if err != nil {
				// Check if the error is due to insufficient permissions
				if strings.Contains(err.Error(), "403 Resource not accessible by integration") {
					fmt.Printf("Error creating release: Insufficient permissions. Please ensure that you have the necessary access rights.\n")
					fmt.Printf("To fix this, go to https://github.com/%s/%s/settings/actions, under \"Workflow permissions\" set \"Read and write permissions\".\n", repoOwner, repoName)
				} else {
					logError(fmt.Sprintf("Error creating release: %v", err))
				}
			} else {
				logInfo("New release created successfully!")
				logVerbose(fmt.Sprintf("Release ID: %v", *createdRelease.ID))
			}
		} else {
			// Another error occurred while retrieving the release
			logError(fmt.Sprintf("Error retrieving release: %v", err))
		}
	} else {
		// The release exists, compare the commit hashes
		logVerbose(fmt.Sprintf("Release found with ID: %d", *release.ID))
		if *release.TargetCommitish != releaseCommitHash {
			logVerbose("Existing release commit hash differs from the desired one. Deleting the existing release and tag...")
			_, err := client.Repositories.DeleteRelease(ctx, repoOwner, repoName, *release.ID)
			if err != nil {
				logError(fmt.Sprintf("Error deleting release: %v", err))
			} else {
				logInfo("Existing release deleted successfully.")
				_, err := client.Git.DeleteRef(ctx, repoOwner, repoName, fmt.Sprintf("tags/%s", releaseTag))
				if err != nil {
					logError(fmt.Sprintf("Error deleting tag: %v", err))
				} else {
					logInfo("Existing tag deleted successfully.")
				}

				// Proceed to create a new release to replace the deleted one
				newRelease := &github.RepositoryRelease{
					TagName:         &releaseTag,
					TargetCommitish: &releaseCommitHash,
					Name:            &releaseName,
				}
				createdRelease, _, err := client.Repositories.CreateRelease(ctx, repoOwner, repoName, newRelease)
				if err != nil {
					logError(fmt.Sprintf("Error creating release: %v", err))
				} else {
					logInfo("New release created successfully!")
					logVerbose(fmt.Sprintf("Release ID: %v", *createdRelease.ID))
				}
			}
		} else {
			logInfo("Release with the name 'continuous' already exists and has the desired commit hash.")
		}
	}
}

func extractRepositoryName(fullName string) string {
	parts := strings.Split(fullName, "/")
	if len(parts) > 1 {
		return parts[1]
	}
	return fullName
}

func logInfo(msg string) {
	fmt.Printf("[INFO] %s\n", msg)
}

func logVerbose(msg string) {
	fmt.Printf("[VERBOSE] %s\n", msg)
}

func logError(msg string) {
	fmt.Printf("[ERROR] %s\n", msg)
}
