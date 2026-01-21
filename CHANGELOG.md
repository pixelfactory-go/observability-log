# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.4.0](https://github.com/pixelfactory-go/observability-log/compare/v1.3.2...v1.4.0) (2026-01-21)


### Features

* add comprehensive OSS-Fuzz integration with native Go fuzzing ([#16](https://github.com/pixelfactory-go/observability-log/issues/16)) ([e2e011e](https://github.com/pixelfactory-go/observability-log/commit/e2e011ec1adbd4adda2776ba5d4e75e08168c76e))

## [1.3.2](https://github.com/pixelfactory-go/observability-log/compare/v1.3.1...v1.3.2) (2026-01-21)


### Bug Fixes

* add name to test job in CI configuration ([#14](https://github.com/pixelfactory-go/observability-log/issues/14)) ([adcf7c6](https://github.com/pixelfactory-go/observability-log/commit/adcf7c6e8dc998cd4e8e2f6bc21af4b7d1c6b541))

## [1.3.1](https://github.com/pixelfactory-go/observability-log/compare/v1.3.0...v1.3.1) (2026-01-19)


### Bug Fixes

* remove deprecated package-name parameter from release-please ([#9](https://github.com/pixelfactory-go/observability-log/issues/9)) ([cf27170](https://github.com/pixelfactory-go/observability-log/commit/cf271700e404940604ab75b2f322aaccb944d539))

## [Unreleased]

### Added
- Comprehensive documentation including README, CONTRIBUTING, and CHANGELOG
- CI/CD workflow consolidation with test matrix for Go 1.24.x and 1.25.x
- Release automation with release-please
- Enhanced Makefile with build and help targets
- Code quality improvements with golangci-lint v2 configuration

### Changed
- Updated golangci-lint configuration to v2 format with comprehensive linter settings
- Updated Go version support to 1.23 minimum
- Updated all dependencies to latest versions
- Improved CI/CD workflows for better efficiency

## [0.3.0] - 2024-10-01

### Changed
- Updated ecszap to version v0.3.0
- Dependency updates for improved compatibility

## [0.2.0] - 2023-05-15

### Changed
- Return struct instead of interface for better type safety
- Updated to Go version 1.23

## [0.1.0] - 2020-07-31

### Added
- Initial release with ECS v1.5.0 support
- Zap-based structured logging
- Sentry integration for error tracking
- Field helpers for HTTP requests/responses
- User agent parsing
- Service and source field helpers

[Unreleased]: https://github.com/pixelfactory-go/observability-log/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/pixelfactory-go/observability-log/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/pixelfactory-go/observability-log/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/pixelfactory-go/observability-log/releases/tag/v0.1.0
