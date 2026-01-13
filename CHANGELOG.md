# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.0.0] - 2026-01-13

This is the first stable release of the vectors library. The library provides comprehensive vector mathematics operations with full support for both float32 and float64 through Go generics, backward compatibility with existing code, and extensive BDD test coverage.

### Added

- Generic type support: `Vec[T Float]` supporting both `float32` and `float64` ([#10](https://github.com/SPANDigital/vectors/pull/10))
- Type alias `Vector = Vec[float64]` for backward compatibility
- Type-specific epsilon constants: `DefaultEpsilon64` (1e-9) and `DefaultEpsilon32` (1e-7)
- Vector arithmetic operations: `Add()` and `Subtract()` methods ([#9](https://github.com/SPANDigital/vectors/pull/9))
- `DotProduct()` function for standard dot product calculation ([#8](https://github.com/SPANDigital/vectors/pull/8))
- `EqualsWithEpsilon()` method for epsilon-based vector equality comparison ([#6](https://github.com/SPANDigital/vectors/pull/6))
- Distance metrics: `EuclideanDistance()`, `TaxicabDistance()`, and `NegativeInnerProduct()` ([#5](https://github.com/SPANDigital/vectors/pull/5))
- `Magnitude()` method to calculate vector length ([#4](https://github.com/SPANDigital/vectors/pull/4))
- Comprehensive documentation in README.md with examples for all operations
- Full BDD test coverage using Cucumber/godog with feature files
- Float32-specific test suite (`vectors_float32_test.go`)
- Development guidelines in CLAUDE.md

### Changed

- All operations now preserve type (float32 or float64) through generic implementation
- `DefaultEpsilon` deprecated in favor of type-specific `DefaultEpsilon64` and `DefaultEpsilon32`
- Enhanced documentation with type usage examples and best practices
- Updated `CosineSimilarity()`, `Equals()`, `Normalize()`, and `IsNormalized()` to use generic types

### Dependencies

- Updated github.com/cucumber/godog from 0.15.0 to 0.15.1 ([#3](https://github.com/SPANDigital/vectors/pull/3))

## [0.2.1] - 2025-02-19

### Fixed

- Pre-commit hook configuration and test execution
- BDD test infrastructure improvements

### Changed

- Improved test coverage for existing features
- Enhanced test reliability

## [0.2.0] - 2024-07-11

### Added

- `Normalize()` method to convert vectors to unit length
- `IsNormalized()` method to check if a vector has unit length
- Epsilon-based comparison for floating-point precision in `IsNormalized()`
- Zero vector handling: `Normalize()` returns zero vector unchanged

### Changed

- Extended documentation with normalization examples
- Improved test coverage for normalization operations

## [0.1.0] - 2024-07-05

### Added

- Initial release with `CosineSimilarity()` function
- `Vector` type definition ([]float64)
- `Equals()` method for exact vector equality
- Error handling for zero vectors in cosine similarity
- BDD test framework using Cucumber/godog
- Comprehensive project documentation and contributing guidelines
- MIT License
- Dev Containers support for consistent development environment
- GitHub Actions CI workflow for automated testing
- Pre-commit hooks for code quality (go-fmt, golangci-lint, go-unit-tests, go-mod-tidy, commitlint, reformat-gherkin)
- Security policy (SECURITY.md)

[Unreleased]: https://github.com/SPANDigital/vectors/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/SPANDigital/vectors/compare/v0.2.1...v1.0.0
[0.2.1]: https://github.com/SPANDigital/vectors/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/SPANDigital/vectors/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/SPANDigital/vectors/releases/tag/v0.1.0
