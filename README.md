# vectors

A lightweight Go library for vector mathematics operations, designed for machine learning and computational geometry applications.

[![Open in Dev Containers](https://img.shields.io/static/v1?label=Dev%20Containers&message=Open&color=blue&logo=visualstudiocode)](https://vscode.dev/redirect?url=vscode://ms-vscode-remote.remote-containers/cloneInVolume?url=https://github.com/SPANDigital/vectors)
![Dev Go Action Workflow Status](https://img.shields.io/github/actions/workflow/status/spandigital/vectors/go.yml?branch=dev&label=dev)
![Main Go Action Workflow Status](https://img.shields.io/github/actions/workflow/status/spandigital/vectors/go.yml?branch=main&label=main)
![Tag](https://img.shields.io/github/v/tag/SPANDigital/vectors)
[![Go Reference](https://pkg.go.dev/badge/github.com/spandigital/vectors.svg)](https://pkg.go.dev/github.com/spandigital/vectors)

## Features

- **Magnitude**: Calculate the Euclidean length of vectors
- **Normalization**: Convert vectors to unit length
- **Vector Arithmetic**: Add and subtract vectors element-wise
- **Cosine Similarity**: Measure similarity between vectors
- **Euclidean Distance**: Calculate straight-line distance between vectors
- **Negative Inner Product**: Calculate negative dot product as a distance metric
- **Dot Product**: Calculate standard dot product between vectors
- **Taxicab Distance**: Calculate Manhattan/L1 distance between vectors
- **Equality Checking**: Compare vectors for exact or epsilon-based equality
- **Normalization Verification**: Check if a vector is already normalized
- **BDD Test Coverage**: Comprehensive behavioral tests with Cucumber/godog

## Installation

```bash
go get github.com/spandigital/vectors
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/spandigital/vectors"
)

func main() {
    // Create a vector
    v := vectors.Vector{3.0, 4.0}

    // Calculate magnitude
    magnitude := v.Magnitude()
    fmt.Printf("Magnitude: %.2f\n", magnitude) // Output: Magnitude: 5.00

    // Normalize the vector
    normalized := v.Normalize()
    fmt.Printf("Normalized: %v\n", normalized) // Output: Normalized: [0.6 0.8]

    // Check if normalized
    isNorm := normalized.IsNormalized()
    fmt.Printf("Is normalized: %t\n", isNorm) // Output: Is normalized: true

    // Add vectors
    v3 := vectors.Vector{1.0, 2.0}
    sum, err := v.Add(v3)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Sum: %v\n", sum) // Output: Sum: [4 6]

    // Subtract vectors
    diff, err := v.Subtract(v3)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Difference: %v\n", diff) // Output: Difference: [2 2]

    // Calculate cosine similarity
    v2 := vectors.Vector{1.0, 2.0}
    similarity, err := vectors.CosineSimilarity(v, v2)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Cosine similarity: %.4f\n", similarity)

    // Calculate Euclidean distance
    distance, err := vectors.EuclideanDistance(v, v2)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Euclidean distance: %.2f\n", distance)
}
```

## API Reference

### Type Definition

```go
type Vector []float64
```

Vector is an alias for a slice of float64 values, representing an n-dimensional vector.

### Methods

#### Magnitude

```go
func (v Vector) Magnitude() float64
```

Returns the Euclidean length (magnitude) of the vector using the formula: √(Σx²)

**Example:**
```go
v := vectors.Vector{3.0, 4.0}
magnitude := v.Magnitude() // Returns 5.0

v2 := vectors.Vector{1.0, 2.0, 2.0}
magnitude2 := v2.Magnitude() // Returns 3.0
```

#### Normalize

```go
func (v Vector) Normalize() Vector
```

Returns a new vector with the same direction but unit length (magnitude of 1.0). If the vector is a zero vector, it returns the zero vector unchanged.

**Example:**
```go
v := vectors.Vector{3.0, 4.0}
normalized := v.Normalize() // Returns [0.6, 0.8]

zero := vectors.Vector{0.0, 0.0}
normalizedZero := zero.Normalize() // Returns [0.0, 0.0]
```

#### IsNormalized

```go
func (v Vector) IsNormalized() bool
```

Checks whether the vector has unit length (magnitude of 1.0). Uses epsilon comparison (1e-9) for floating-point precision.

**Example:**
```go
v := vectors.Vector{0.6, 0.8}
isNorm := v.IsNormalized() // Returns true

v2 := vectors.Vector{3.0, 4.0}
isNorm2 := v2.IsNormalized() // Returns false
```

#### Equals

```go
func (a Vector) Equals(b Vector) bool
```

Compares two vectors for exact equality. Returns false if vectors have different lengths or any components differ.

**Example:**
```go
v1 := vectors.Vector{1.0, 2.0, 3.0}
v2 := vectors.Vector{1.0, 2.0, 3.0}
v3 := vectors.Vector{1.0, 2.0}

v1.Equals(v2) // Returns true
v1.Equals(v3) // Returns false (different lengths)
```

#### EqualsWithEpsilon

```go
func (v Vector) EqualsWithEpsilon(other Vector, epsilon float64) bool
```

Compares two vectors for equality within a specified tolerance. Returns true if the absolute difference between each corresponding component is less than or equal to epsilon. Returns false if vectors have different lengths.

This method is useful for comparing floating-point vectors where exact equality is impractical due to precision limitations.

**Example:**
```go
v1 := vectors.Vector{1.0, 2.0, 3.0}
v2 := vectors.Vector{1.0001, 2.0001, 3.0001}

// Exact comparison fails
v1.Equals(v2) // Returns false

// Epsilon comparison succeeds
v1.EqualsWithEpsilon(v2, 0.001) // Returns true
v1.EqualsWithEpsilon(v2, 0.00001) // Returns false

// Using the default epsilon constant
v3 := vectors.Vector{1.0, 2.0}
v4 := vectors.Vector{1.000000001, 2.000000001}
v3.EqualsWithEpsilon(v4, vectors.DefaultEpsilon) // Returns true

// Different lengths still return false
v5 := vectors.Vector{1.0, 2.0}
v6 := vectors.Vector{1.0}
v5.EqualsWithEpsilon(v6, 0.1) // Returns false
```

#### Add

```go
func (v Vector) Add(other Vector) (Vector, error)
```

Returns a new vector that is the element-wise sum of v and other. Returns an error if the vectors have different lengths.

The operation computes: result[i] = v[i] + other[i]

**Example:**
```go
v1 := vectors.Vector{1.0, 2.0, 3.0}
v2 := vectors.Vector{4.0, 5.0, 6.0}

sum, err := v1.Add(v2)
if err != nil {
    panic(err)
}
fmt.Printf("Sum: %v\n", sum) // Output: Sum: [5 7 9]

// Add with negative components
v3 := vectors.Vector{-1.0, 2.0}
v4 := vectors.Vector{3.0, -4.0}
sum2, _ := v3.Add(v4) // Returns [2.0, -2.0]

// Identity: adding zero vector
v5 := vectors.Vector{3.0, 4.0}
zero := vectors.Vector{0.0, 0.0}
sum3, _ := v5.Add(zero) // Returns [3.0, 4.0]

// Different lengths return error
v6 := vectors.Vector{1.0, 2.0}
v7 := vectors.Vector{1.0}
_, err = v6.Add(v7) // Returns error
```

#### Subtract

```go
func (v Vector) Subtract(other Vector) (Vector, error)
```

Returns a new vector that is the element-wise difference of v and other. Returns an error if the vectors have different lengths.

The operation computes: result[i] = v[i] - other[i]

**Example:**
```go
v1 := vectors.Vector{5.0, 7.0, 9.0}
v2 := vectors.Vector{1.0, 2.0, 3.0}

diff, err := v1.Subtract(v2)
if err != nil {
    panic(err)
}
fmt.Printf("Difference: %v\n", diff) // Output: Difference: [4 5 6]

// Subtract with negative components
v3 := vectors.Vector{3.0, -4.0}
v4 := vectors.Vector{-1.0, 2.0}
diff2, _ := v3.Subtract(v4) // Returns [4.0, -6.0]

// Identity: subtracting zero vector
v5 := vectors.Vector{3.0, 4.0}
zero := vectors.Vector{0.0, 0.0}
diff3, _ := v5.Subtract(zero) // Returns [3.0, 4.0]

// Inverse: subtracting from itself returns zero vector
v6 := vectors.Vector{3.0, 4.0}
diff4, _ := v6.Subtract(v6) // Returns [0.0, 0.0]

// Different lengths return error
v7 := vectors.Vector{1.0, 2.0, 3.0}
v8 := vectors.Vector{1.0, 2.0}
_, err = v7.Subtract(v8) // Returns error
```

### Constants

```go
const DefaultEpsilon = 1e-9
```

The default epsilon value used for floating-point comparisons throughout the library. This constant is used internally by `IsNormalized()` and can be used as a standard tolerance for `EqualsWithEpsilon()`.

### Functions

#### CosineSimilarity

```go
func CosineSimilarity(a Vector, b Vector) (float64, error)
```

Calculates the cosine similarity between two vectors. Returns a value between -1.0 and 1.0, where:
- 1.0 indicates identical direction
- 0.0 indicates orthogonal vectors
- -1.0 indicates opposite direction

Returns an error if either vector is a zero vector.

**Example:**
```go
v1 := vectors.Vector{1.0, 2.0, 3.0}
v2 := vectors.Vector{4.0, 5.0, 6.0}

similarity, err := vectors.CosineSimilarity(v1, v2)
if err != nil {
    // Handle error (e.g., zero vector)
    panic(err)
}
fmt.Printf("Similarity: %.4f\n", similarity)

// Identical vectors
v3 := vectors.Vector{1.0, 0.0}
similarity2, _ := vectors.CosineSimilarity(v3, v3) // Returns 1.0

// Orthogonal vectors
v4 := vectors.Vector{1.0, 0.0}
v5 := vectors.Vector{0.0, 1.0}
similarity3, _ := vectors.CosineSimilarity(v4, v5) // Returns 0.0
```

#### EuclideanDistance

```go
func EuclideanDistance(a Vector, b Vector) (float64, error)
```

Calculates the Euclidean distance (straight-line distance) between two vectors using the formula: √(Σ(ai - bi)²)

Returns an error if the vectors have different lengths.

**Example:**
```go
v1 := vectors.Vector{1.0, 2.0}
v2 := vectors.Vector{4.0, 6.0}

distance, err := vectors.EuclideanDistance(v1, v2)
if err != nil {
    panic(err)
}
fmt.Printf("Distance: %.1f\n", distance) // Returns 5.0

// Identical vectors have zero distance
v3 := vectors.Vector{1.0, 2.0, 3.0}
distance2, _ := vectors.EuclideanDistance(v3, v3) // Returns 0.0

// Distance from origin
origin := vectors.Vector{0.0, 0.0, 0.0}
point := vectors.Vector{1.0, 2.0, 2.0}
distance3, _ := vectors.EuclideanDistance(origin, point) // Returns 3.0
```

#### NegativeInnerProduct

```go
func NegativeInnerProduct(a Vector, b Vector) (float64, error)
```

Calculates the negative inner product (negative dot product) between two vectors using the formula: -(Σ(ai * bi))

This is commonly used as a distance metric in machine learning and information retrieval. Returns an error if the vectors have different lengths.

**Example:**
```go
v1 := vectors.Vector{1.0, 2.0, 3.0}
v2 := vectors.Vector{4.0, 5.0, 6.0}

product, err := vectors.NegativeInnerProduct(v1, v2)
if err != nil {
    panic(err)
}
fmt.Printf("Negative inner product: %.1f\n", product) // Returns -32.0

// Orthogonal vectors have zero inner product
v3 := vectors.Vector{1.0, 0.0}
v4 := vectors.Vector{0.0, 1.0}
product2, _ := vectors.NegativeInnerProduct(v3, v4) // Returns 0.0

// With negative components
v5 := vectors.Vector{-1.0, 2.0}
v6 := vectors.Vector{3.0, -4.0}
product3, _ := vectors.NegativeInnerProduct(v5, v6) // Returns 11.0
```

**Note**: If you need the standard (positive) dot product, use `DotProduct()` instead. `NegativeInnerProduct` is specifically designed for use as a distance metric in machine learning applications.

#### DotProduct

```go
func DotProduct(a Vector, b Vector) (float64, error)
```

Calculates the dot product (inner product) between two vectors using the formula: Σ(ai * bi)

The dot product is a fundamental vector operation used for:
- Computing projections
- Measuring similarity (when normalized)
- Calculating angles between vectors
- Many machine learning operations

Returns an error if the vectors have different lengths.

**Example:**
```go
v1 := vectors.Vector{1.0, 2.0, 3.0}
v2 := vectors.Vector{4.0, 5.0, 6.0}

product, err := vectors.DotProduct(v1, v2)
if err != nil {
    panic(err)
}
fmt.Printf("Dot product: %.1f\n", product) // Returns 32.0

// Orthogonal vectors have zero dot product
v3 := vectors.Vector{1.0, 0.0}
v4 := vectors.Vector{0.0, 1.0}
product2, _ := vectors.DotProduct(v3, v4) // Returns 0.0

// With negative components
v5 := vectors.Vector{-1.0, 2.0}
v6 := vectors.Vector{3.0, -4.0}
product3, _ := vectors.DotProduct(v5, v6) // Returns -11.0

// Different lengths return error
v7 := vectors.Vector{1.0, 2.0}
v8 := vectors.Vector{1.0}
_, err = vectors.DotProduct(v7, v8) // Returns error
```

#### TaxicabDistance

```go
func TaxicabDistance(a Vector, b Vector) (float64, error)
```

Calculates the taxicab distance (Manhattan distance, L1 distance) between two vectors using the formula: Σ|ai - bi|

This metric represents the distance when only axis-aligned movement is allowed. Returns an error if the vectors have different lengths.

**Example:**
```go
v1 := vectors.Vector{1.0, 2.0}
v2 := vectors.Vector{4.0, 6.0}

distance, err := vectors.TaxicabDistance(v1, v2)
if err != nil {
    panic(err)
}
fmt.Printf("Taxicab distance: %.1f\n", distance) // Returns 7.0

// Distance from origin
origin := vectors.Vector{0.0, 0.0, 0.0}
point := vectors.Vector{1.0, 2.0, 3.0}
distance2, _ := vectors.TaxicabDistance(origin, point) // Returns 6.0

// With negative components
v3 := vectors.Vector{-1.0, -2.0}
v4 := vectors.Vector{2.0, 3.0}
distance3, _ := vectors.TaxicabDistance(v3, v4) // Returns 8.0
```

## Development

### Prerequisites

- Go 1.22.2 or later
- Pre-commit hooks (optional but recommended)

### Setup

```bash
# Clone the repository
git clone https://github.com/SPANDigital/vectors.git
cd vectors

# Install dependencies
go mod download

# Install pre-commit hooks (optional)
pre-commit install
```

### Running Tests

```bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -v -cover ./...
```

The project uses Behavior-Driven Development (BDD) with Cucumber/godog. Feature files are located in the `features/` directory.

### Building

```bash
go build -v ./...
```

### Code Quality

This project uses pre-commit hooks to ensure code quality:
- **go-fmt**: Automatic code formatting
- **golangci-lint**: Comprehensive linting
- **go-unit-tests**: Automatic test execution
- **go-mod-tidy**: Dependency management
- **commitlint**: Conventional commit message validation
- **reformat-gherkin**: Feature file formatting

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details on:
- Code of Conduct
- Development workflow
- Pull request process
- Testing requirements

### Commit Messages

This project follows [Conventional Commits](https://www.conventionalcommits.org/). Format your commits as:

```
<type>: <description>

[optional body]

[optional footer]
```

Types: `feat`, `fix`, `docs`, `test`, `chore`, `refactor`, `style`, `perf`

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/SPANDigital/vectors/issues)
- **Security**: See [SECURITY.md](SECURITY.md) for reporting vulnerabilities
- **Contact**: richard.wooding@spandigital.com

## Acknowledgments

Developed by [Span Digital](https://github.com/SPANDigital) with comprehensive test coverage using Cucumber/godog for behavior-driven development.
