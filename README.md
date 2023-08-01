# Continuous Release Manager

This tool provides a "continuous" release that always reflects the latest commit hash of the branch being built on CI (e.g., **GitHub Actions** or **Cirrus CI**). The tool checks for an existing release named "continuous" in the repository. If the release does not exist, it creates a new one, or if it exists and its commit hash differs, it deletes the existing release and creates a new one with the latest commit hash. This ensures that the "continuous" release always points to the correct hash. The tool can be easily integrated into your repository's CI/CD pipeline, providing a reliable and automated way to manage continuous releases. It is especially useful if you want to upload binaries to releases from different CI systems (e.g., GitHub Actions and Cirrus CI).

## How it Works

1. The tool checks if a release with the name "continuous" already exists.
   - If a release with the name "continuous" exists, it compares the commit hash of the existing release with the desired commit hash.
   - If the commit hashes differ, the existing release is deleted to keep the releases in sync with the current state of the code.

2. After checking for an existing release, the tool creates a new release named "continuous" with the desired commit hash.

## Usage

To use the tool, set up a GitHub Actions workflow (or any CI/CD system) to automatically trigger the tool whenever changes are pushed to the repository. The tool will handle creating, updating, or deleting the "continuous" release as needed.

## Examples

## GitHub Actions (for the Linux build)

```yaml
      - name: Create GitHub Release using Continuous Release Manager
        if: github.event_name == 'push'  # Only run for push events, not pull requests
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        run: |
          curl -L -o continuous-release-manager-linux https://github.com/probonopd/continuous-release-manager/releases/download/continuous/continuous-release-manager-linux
          chmod +x continuous-release-manager-linux
          ./continuous-release-manager-linux

      - name: Upload to GitHub Release
        if: github.event_name == 'push'  # Only run for push events, not pull requests
        uses: xresloader/upload-to-github-release@v1.3.12
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          file: "continuous-release-manager-*"
          draft: false
          verbose: true
          branches: main
          update_latest_release: true
          tag_name: continuous
```

## Cirrus CI (for the FreeBSD build)

```yaml
task:
  ...
  test_script:
    - ...
    - case "$CIRRUS_BRANCH" in *pull/*) echo "Skipping since it's a pull request" ;; * ) wget https://github.com/tcnksm/ghr/files/5247714/ghr.zip ; unzip ghr.zip ; rm ghr.zip ; fetch https://github.com/probonopd/continuous-release-manager/releases/download/continuous/continuous-release-manager-freebsd && chmod +x continuous-release-manager-freebsd && ./continuous-release-manager-freebsd && ./ghr -replace -t "${GITHUB_TOKEN}" -u "${CIRRUS_REPO_OWNER}" -r "${CIRRUS_REPO_NAME}" -c "${CIRRUS_CHANGE_IN_REPO}" continuous "${CIRRUS_WORKING_DIR}"/build/*zip ; esac
  only_if: $CIRRUS_TAG != 'continuous'
```

## Example output

### For the first builder

```
[INFO] Starting release management...
[VERBOSE] Repository Owner: probonopd
[VERBOSE] Repository Name: Filer
[VERBOSE] Release Tag: continuous
[VERBOSE] Release Commit Hash: d8b990a61fcb2671a8621d282341fdcbf1e83cc7
[INFO] Checking for existing release...
[VERBOSE] Release found with ID: 114646674
[VERBOSE] Existing release commit hash differs from the desired one. Deleting the existing release...
[INFO] Existing release deleted successfully.
[INFO] New release created successfully!
[VERBOSE] Release ID: 114646855
```

### For the subsequent builders

```
[INFO] Starting release management...
[VERBOSE] Repository Owner: probonopd
[VERBOSE] Repository Name: Filer
[VERBOSE] Release Tag: continuous
[VERBOSE] Release Commit Hash: d8b990a61fcb2671a8621d282341fdcbf1e83cc7
[INFO] Checking for existing release...
[VERBOSE] Release found with ID: 114646855
[INFO] Release with the name 'continuous' already exists and has the desired commit hash.
WARNING: found release (continuous). Use existing one.
```

### Projects using this

* https://github.com/helloSystem/Menu
* https://github.com/helloSystem/launch
* https://github.com/probonopd/Filer
* ...
  
## License

The GitHub Continuous Release Manager is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
_This project is not affiliated with GitHub or any other third-party service mentioned._
