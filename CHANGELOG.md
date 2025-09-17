# Changelog

All notable changes to this project will be documented in this file.

## [v0.5.100] - 2025-09-17

### Removed

- Removed AWS SDK dependency from `go.mod` and `go.sum`.

### Refactored

- Changed all model fields from pointer types to direct value types with appropriate zero values.
- Updated JSON marshaling to use `OmitZeroStructFields` for zero value serialization.
- Removed AWS SDK usage related to string pointers.

Affected files:

- `cmd/example/main.go`
- `reptiloid/models/image/image.go`
- `reptiloid/models/image/image_test.go`
- `reptiloid/models/text/text.go`
- `reptiloid/models/text/text_test.go`
- `reptiloid/service.go`

## [v0.4.102] - 2025-09-09

### Fix

- Capitalize exported types and methods in reptiloid package.

## [v0.4.101] - 2025-09-06

### Changed

- Use pointers for models and inputs in client API and examples.

## [v0.4.100] - 2025-09-06

### Changed

- Updated Client API to use pointers for models and inputs to improve consistency and efficiency.
- Modified NewClient, Generate, and related functions to accept pointers.
- Updated associated tests and example usage in main.go to reflect these changes.

This change streamlines memory usage and standardizes function signatures across the client API.

## [v0.0.0] - 2025-01-01

### Added

- Initial placeholder for changelog entries.
