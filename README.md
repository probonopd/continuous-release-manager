# GitHub Continuous Release Manager

The GitHub Continuous Release Manager is a Go tool that automates the process of managing continuous releases in a GitHub repository. The tool checks if a release with the name "continuous" already exists and compares its commit hash. If the commit hashes differ, the existing release is deleted to avoid conflicts. After verification, the tool creates a new release named "continuous" with the desired commit hash.

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
