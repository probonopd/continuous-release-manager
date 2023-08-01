# GitHub Continuous Release Manager

This tool provides a "continuous" release that always reflects the latest commit hash of the branch being built on CI (e.g., GitHub Actions or Cirrus CI). The tool checks for an existing release named "continuous" in the repository. If the release does not exist, it creates a new one, or if it exists and its commit hash differs, it deletes the existing release and creates a new one with the latest commit hash. This ensures that the "continuous" release always points to the correct hash. The tool can be easily integrated into your repository's CI/CD pipeline, providing a reliable and automated way to manage continuous releases. It is especially useful if you want to upload binaries to releases from different CI systems (e.g., GitHub Actions and Cirrus CI).

## How it Works

1. The tool checks if a release with the name "continuous" already exists.
   - If a release with the name "continuous" exists, it compares the commit hash of the existing release with the desired commit hash.
   - If the commit hashes differ, the existing release is deleted to keep the releases in sync with the current state of the code.

2. After checking for an existing release, the tool creates a new release named "continuous" with the desired commit hash.

## Usage

To use the tool, set up a GitHub Actions workflow (or any CI/CD system) to automatically trigger the tool whenever changes are pushed to the repository. The tool will handle creating, updating, or deleting the "continuous" release as needed.

## GitHub Actions Example

```yaml
...
```

## License

The GitHub Continuous Release Manager is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
_This project is not affiliated with GitHub or any other third-party service mentioned._
