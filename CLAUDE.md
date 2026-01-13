# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Building and Testing
```bash
# Build the project
go build -v ./...

# Run all tests (includes BDD feature tests)
go test -v ./...
```

### Pre-commit Hooks
This project uses pre-commit hooks. After making changes, the following automatically run:
- `go-fmt` - Format Go code
- `golangci-lint` - Lint Go code
- `go-unit-tests` - Run unit tests
- `go-mod-tidy` - Tidy Go modules
- `commitlint` - Enforce conventional commits
- `reformat-gherkin` - Format .feature files

To manually run pre-commit on all files:
```bash
pre-commit run --all-files
```

## Architecture Overview

### Core Design
This is a Go library providing vector mathematics operations. The architecture follows these principles:

- **Generic Type Definition**: `Vec[T Float] []T` - vectors support both float32 and float64
- **Backward Compatibility**: `Vector = Vec[float64]` - type alias ensures existing code works unchanged
- **Receiver Methods**: Operations implemented as generic methods on the Vec[T] type (e.g., `v.Magnitude()`, `v.Normalize()`)
- **Type Preservation**: All operations return the same type as their inputs (e.g., `Vec[float32].Magnitude()` returns `float32`)
- **Error Handling**: Operations that can fail return `(result, error)` tuples (e.g., `CosineSimilarity` returns error for zero vectors)
- **BDD Testing**: Comprehensive test coverage using Cucumber/godog with Gherkin feature files

### Working with Generic Types

When adding new operations, follow these patterns:

**Methods** (operate on a single vector or return vector results):
```go
func (v Vec[T]) NewMethod() T {
    sum := T(0)  // Use T(0) instead of 0.0
    for _, component := range v {
        sum += component
    }
    return sum
}
```

**Functions** (operate on multiple vectors):
```go
func NewFunction[T Float](a Vec[T], b Vec[T]) (T, error) {
    if len(a) != len(b) {
        return T(0), errors.New("vectors must have the same length")
    }
    // implementation
}
```

**Math Operations** (convert to float64, then back to T):
```go
func (v Vec[T]) OperationWithMath() T {
    value := v.someValue()
    result := T(math.Sqrt(float64(value)))  // Always convert for math ops
    return result
}
```

**Type-Specific Epsilon**:
```go
func (v Vec[T]) IsSpecialProperty() bool {
    eps := epsilon[T]()  // Gets DefaultEpsilon32 for float32, DefaultEpsilon64 for float64
    // use eps for comparisons
}
```

**Creating New Vectors**:
```go
result := make(Vec[T], len(v))  // Use Vec[T], not Vector
```

### File Organization Pattern
Each vector operation follows this structure:
- `operation.go` - Implementation file with the method/function
- `operation_test.go` - BDD test step definitions and scenario initialization
- `features/operation.feature` - Gherkin feature file with test scenarios

Example for the Magnitude operation:
- `magnitude.go` - Contains `func (v Vec[T]) Magnitude() T`
- `magnitude_test.go` - Contains `iCalculateMagnitudeOfVector()` and `initializeMagnitudeScenario()`
- `features/magnitude.feature` - Contains BDD test scenarios

### BDD Testing Structure
Tests use a shared context pattern defined in `vectors_test.go`:

```go
type aKey struct{}      // Context key for first vector
type bKey struct{}      // Context key for second vector
type resultKey struct{} // Context key for operation result
type errKey struct{}    // Context key for errors
```

All scenario initializers must be registered in `InitializeScenario()` in `vectors_test.go`.

Shared helper functions:
- `convertTableToVector()` - Converts Gherkin tables to Vector type
- `vectorIs()` - Step definition factory for "Given vector X is" steps
- `theResultShouldBe()` - Assertion for float64 results
- `theResultShouldBeVector()` - Assertion for Vector results
- `theResultShouldBeBool()` - Assertion for boolean results
- `theResultShouldBeAnError()` - Assertion for error results

## Adding New Vector Operations

When adding a new operation, follow this pattern:

1. **Create implementation file** (e.g., `newoperation.go`):
   ```go
   package vectors

   func (v Vector) NewOperation() ResultType {
       // implementation
   }
   ```

2. **Create test file** (e.g., `newoperation_test.go`):
   ```go
   package vectors

   import (
       "context"
       "errors"
       "github.com/cucumber/godog"
   )

   func iPerformNewOperation(key any) func(ctx context.Context) (context.Context, error) {
       return func(ctx context.Context) (context.Context, error) {
           v, ok := ctx.Value(key).(Vector)
           if !ok {
               return ctx, errors.New("vector not found in context")
           }
           return context.WithValue(ctx, resultKey{}, v.NewOperation()), nil
       }
   }

   func initializeNewOperationScenario(ctx *godog.ScenarioContext) {
       ctx.Step(`^I perform new operation on vector a$`, iPerformNewOperation(aKey{}))
   }
   ```

3. **Create feature file** (`features/newoperation.feature`):
   ```gherkin
   Feature: New Operation
     As a developer
     I want to perform new operation on vectors
     So that I can achieve specific goal

     Scenario: Description of test case
       Given vector a is
         | 1.0 |
         | 2.0 |
       When I perform new operation on vector a
       Then the result should be <expected>
   ```

4. **Register the scenario initializer** in `vectors_test.go`:
   ```go
   func InitializeScenario(ctx *godog.ScenarioContext) {
       initializeSharedSteps(ctx)
       // ... existing initializers ...
       initializeNewOperationScenario(ctx) // Add this line
   }
   ```

## Special Considerations

### Floating Point Precision
When comparing floating point values for equality (e.g., checking if a vector is normalized), use epsilon comparison:
```go
const epsilon = 1e-9
if math.Abs(value - expected) < epsilon { ... }
```

### Zero Vector Handling
- `Normalize()` returns the zero vector unchanged (no error)
- `CosineSimilarity()` returns an error for zero vectors
- Consider edge cases carefully when implementing new operations

### Commit Message Format
This project uses Conventional Commits. Format your commits as:
```
<type>: <description>

[optional body]

Co-Authored-By: Claude Sonnet 4.5 <noreply@anthropic.com>
```

Types: `feat`, `fix`, `chore`, `docs`, `test`, `refactor`, etc.

### Pull Request Requirements
Before creating a PR, ensure:
- Branch has been rebased with the target branch (usually `dev`)
- TDD has been applied with BDD feature tests
- All linters pass
- Documentation added to godoc and/or README.md
- Pre-commit hooks pass
