package vectors

// Float is a constraint that permits float32 and float64 types,
// including named types based on these underlying types.
type Float interface {
	~float32 | ~float64
}

// Vec is a generic vector type representing an n-dimensional vector
// with components of type T, where T must be a floating-point type.
//
// Vec supports both float32 and float64, allowing users to choose
// between precision (float64) and performance/memory efficiency (float32).
//
// All operations on Vec preserve the type of their inputs. For example,
// Vec[float32].Magnitude() returns float32, not float64.
//
// Example:
//
//	v64 := vectors.Vec[float64]{3.0, 4.0}
//	mag64 := v64.Magnitude() // Returns float64: 5.0
//
//	v32 := vectors.Vec[float32]{3.0, 4.0}
//	mag32 := v32.Magnitude() // Returns float32: 5.0
type Vec[T Float] []T

// Vector is a type alias for backward compatibility with existing code.
// It represents a vector with float64 components.
//
// New code may use Vec[float64] directly or continue using Vector.
// They are identical types and fully interchangeable.
type Vector = Vec[float64]

// Epsilon constants for floating-point comparisons.
const (
	// DefaultEpsilon64 is the default tolerance for float64 comparisons.
	// This represents a very small value suitable for 64-bit floating-point precision.
	DefaultEpsilon64 = 1e-9

	// DefaultEpsilon32 is the default tolerance for float32 comparisons.
	// This represents a small value suitable for 32-bit floating-point precision.
	DefaultEpsilon32 = 1e-7

	// DefaultEpsilon is the legacy epsilon constant for backward compatibility.
	// New code should use DefaultEpsilon64 or DefaultEpsilon32 explicitly.
	//
	// Deprecated: Use DefaultEpsilon64 or DefaultEpsilon32 instead.
	DefaultEpsilon = DefaultEpsilon64
)

// epsilon returns the appropriate default epsilon for type T.
// This is a helper function for internal use by operations that need
// type-specific epsilon values.
func epsilon[T Float]() T {
	var zero T
	switch any(zero).(type) {
	case float32:
		return T(DefaultEpsilon32)
	default: // float64
		return T(DefaultEpsilon64)
	}
}
