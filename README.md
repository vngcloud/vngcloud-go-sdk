# VngCloud Go SDK

[![Release VngCloud GoSDK project](https://github.com/vngcloud/vngcloud-go-sdk/actions/workflows/release_build.yml/badge.svg)](https://github.com/vngcloud/vngcloud-go-sdk/actions/workflows/release_build.yml)

<hr>

# Automation with GitHub Actions
- This project uses GitHub action to trigger the build and release processes, please commit with new tags if you are going to release a new version. For example:
  ```bash
  git add .
  git commit -am "Release version v1.0.0"
  git tag v1.0.0
  git push --tags
  ```