# GitHub Continuous Release Manager

The GitHub Continuous Release Manager is a Go tool that automates the process of managing continuous releases in a GitHub repository. The tool checks if a release with the name "continuous" already exists and compares its commit hash. If the commit hashes differ, the existing release is deleted to avoid conflicts. After verification, the tool creates a new release named "continuous" with the desired commit hash.

## How it Works

1. The tool checks if a release with the name "continuous" already exists.
   - If a release with the name "continuous" exists, it compares the commit hash of the existing release with the desired commit hash.
   - If the commit hashes differ, the existing release is deleted to keep the releases in sync with the current state of the code.

2. After checking for an existing release, the tool creates a new release named "continuous" with the desired commit hash.

## Usage

To use the tool, set up a GitHub Actions workflow (or any CI/CD system) to automatically trigger the tool whenever changes are pushed to the repository. The tool will handle creating, updating, or deleting the "continuous" release as needed.

## Installation

1. Clone this repository or copy the "continuous_release_manager.go" file to your project.

2. Set up the required environment variables:
   - `GITHUB_TOKEN`: GitHub personal access token with the necessary permissions to manage releases in your repository.

3. Customize the `repoOwner`, `repoName`, `releaseTag`, and `releaseCommitHash` variables in the tool to match your repository details and desired release information.

4. Run the tool manually or set up a CI/CD system to automate the process.

## GitHub Actions Example

```yaml
name: Build and Continuous Release

on:
  push:
    branches:
      - main

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v2

    - name: Set up Go
      uses: actions/setup-go@v2
      with:
        go-version: '1.16'

    - name: Build binary
      run: go build -o continuous-release-manager

    - name: Create "continuous" release
      run: ./continuous-release-manager
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}

    - name: Upload binary to release
      run: |
        gh release upload continuous ./continuous-release-manager
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

## License

The GitHub Continuous Release Manager is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
_This project is not affiliated with GitHub or any other third-party service mentioned._
