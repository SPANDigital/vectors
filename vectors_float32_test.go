package vectors

import (
	"math"
	"testing"
)

// TestVecFloat32Operations tests basic operations with float32 vectors
func TestVecFloat32Operations(t *testing.T) {
	tests := []struct {
		name string
		op   func() float32
		want float32
	}{
		{
			name: "magnitude of 3-4 vector",
			op: func() float32 {
				v := Vec[float32]{3.0, 4.0}
				return v.Magnitude()
			},
			want: 5.0,
		},
		{
			name: "dot product",
			op: func() float32 {
				v1 := Vec[float32]{1.0, 2.0}
				v2 := Vec[float32]{3.0, 4.0}
				result, _ := DotProduct(v1, v2)
				return result
			},
			want: 11.0,
		},
		{
			name: "euclidean distance",
			op: func() float32 {
				v1 := Vec[float32]{1.0, 2.0}
				v2 := Vec[float32]{4.0, 6.0}
				result, _ := EuclideanDistance(v1, v2)
				return result
			},
			want: 5.0,
		},
		{
			name: "normalize magnitude",
			op: func() float32 {
				v := Vec[float32]{3.0, 4.0}
				normalized := v.Normalize()
				return normalized.Magnitude()
			},
			want: 1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.op()
			if math.Abs(float64(got-tt.want)) > DefaultEpsilon32 {
				t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

// TestVecFloat32VectorOperations tests operations that return vectors
func TestVecFloat32VectorOperations(t *testing.T) {
	v1 := Vec[float32]{1.0, 2.0, 3.0}
	v2 := Vec[float32]{4.0, 5.0, 6.0}

	// Test Add
	sum, err := v1.Add(v2)
	if err != nil {
		t.Fatalf("Add failed: %v", err)
	}
	expected := Vec[float32]{5.0, 7.0, 9.0}
	for i, val := range sum {
		if math.Abs(float64(val-expected[i])) > DefaultEpsilon32 {
			t.Errorf("Add[%d] = %v, want %v", i, val, expected[i])
		}
	}

	// Test Subtract
	diff, err := v2.Subtract(v1)
	if err != nil {
		t.Fatalf("Subtract failed: %v", err)
	}
	expectedDiff := Vec[float32]{3.0, 3.0, 3.0}
	for i, val := range diff {
		if math.Abs(float64(val-expectedDiff[i])) > DefaultEpsilon32 {
			t.Errorf("Subtract[%d] = %v, want %v", i, val, expectedDiff[i])
		}
	}

	// Test Normalize
	v3 := Vec[float32]{3.0, 4.0}
	normalized := v3.Normalize()
	expectedNorm := Vec[float32]{0.6, 0.8}
	for i, val := range normalized {
		if math.Abs(float64(val-expectedNorm[i])) > DefaultEpsilon32 {
			t.Errorf("Normalize[%d] = %v, want %v", i, val, expectedNorm[i])
		}
	}
}

// TestFloat32IsNormalized tests IsNormalized with float32
func TestFloat32IsNormalized(t *testing.T) {
	// Normalized vector
	v1 := Vec[float32]{0.6, 0.8}
	if !v1.IsNormalized() {
		t.Error("Expected Vec[float32]{0.6, 0.8} to be normalized")
	}

	// Non-normalized vector
	v2 := Vec[float32]{3.0, 4.0}
	if v2.IsNormalized() {
		t.Error("Expected Vec[float32]{3.0, 4.0} to not be normalized")
	}

	// Normalized from Normalize()
	v3 := Vec[float32]{3.0, 4.0}
	normalized := v3.Normalize()
	if !normalized.IsNormalized() {
		t.Error("Expected normalized vector to pass IsNormalized()")
	}
}

// TestFloat32Equals tests equality comparisons with float32
func TestFloat32Equals(t *testing.T) {
	v1 := Vec[float32]{1.0, 2.0, 3.0}
	v2 := Vec[float32]{1.0, 2.0, 3.0}
	v3 := Vec[float32]{1.0, 2.0}

	if !v1.Equals(v2) {
		t.Error("Expected identical vectors to be equal")
	}

	if v1.Equals(v3) {
		t.Error("Expected vectors of different lengths to not be equal")
	}

	// Test EqualsWithEpsilon
	v4 := Vec[float32]{1.0, 2.0}
	v5 := Vec[float32]{1.00001, 2.00001}

	if !v4.EqualsWithEpsilon(v5, 0.0001) {
		t.Error("Expected vectors to be equal within epsilon")
	}
}

// TestFloat32EdgeCases tests edge cases specific to float32
func TestFloat32EdgeCases(t *testing.T) {
	// Test very small values don't underflow to zero
	v := Vec[float32]{1e-10, 1e-10}
	mag := v.Magnitude()
	if mag == 0 {
		t.Error("Expected non-zero magnitude for small values")
	}

	// Test zero vector
	zero := Vec[float32]{0.0, 0.0}
	normalized := zero.Normalize()
	if !normalized.Equals(zero) {
		t.Error("Expected zero vector to normalize to itself")
	}
}

// TestFloat32DistanceMetrics tests distance metrics with float32
func TestFloat32DistanceMetrics(t *testing.T) {
	v1 := Vec[float32]{1.0, 2.0}
	v2 := Vec[float32]{4.0, 6.0}

	// Taxicab distance
	taxicab, err := TaxicabDistance(v1, v2)
	if err != nil {
		t.Fatalf("TaxicabDistance failed: %v", err)
	}
	if math.Abs(float64(taxicab-7.0)) > DefaultEpsilon32 {
		t.Errorf("TaxicabDistance = %v, want 7.0", taxicab)
	}

	// Negative inner product
	nip, err := NegativeInnerProduct(v1, v2)
	if err != nil {
		t.Fatalf("NegativeInnerProduct failed: %v", err)
	}
	expectedNip := float32(-16.0) // 1*4 + 2*6 = 4 + 12 = 16, negated = -16
	if math.Abs(float64(nip-expectedNip)) > DefaultEpsilon32 {
		t.Errorf("NegativeInnerProduct = %v, want %v", nip, expectedNip)
	}
}

// TestFloat32CosineSimilarity tests cosine similarity with float32
func TestFloat32CosineSimilarity(t *testing.T) {
	// Identical vectors
	v1 := Vec[float32]{1.0, 2.0, 3.0}
	sim, err := CosineSimilarity(v1, v1)
	if err != nil {
		t.Fatalf("CosineSimilarity failed: %v", err)
	}
	if math.Abs(float64(sim-1.0)) > DefaultEpsilon32 {
		t.Errorf("CosineSimilarity of identical vectors = %v, want 1.0", sim)
	}

	// Orthogonal vectors
	v2 := Vec[float32]{1.0, 0.0}
	v3 := Vec[float32]{0.0, 1.0}
	sim2, err := CosineSimilarity(v2, v3)
	if err != nil {
		t.Fatalf("CosineSimilarity failed: %v", err)
	}
	if math.Abs(float64(sim2)) > DefaultEpsilon32 {
		t.Errorf("CosineSimilarity of orthogonal vectors = %v, want 0.0", sim2)
	}

	// Zero vector should return error
	zero := Vec[float32]{0.0, 0.0}
	v4 := Vec[float32]{1.0, 2.0}
	_, err = CosineSimilarity(zero, v4)
	if err == nil {
		t.Error("Expected error for zero vector in CosineSimilarity")
	}
}

// TestBackwardCompatibility verifies that existing Vector code works unchanged
func TestBackwardCompatibility(t *testing.T) {
	// Test that Vector alias works
	v := Vector{3.0, 4.0}
	mag := v.Magnitude()
	var _ float64 = mag // Should compile - mag is float64

	// Test all existing patterns work
	v2 := Vector{1.0, 2.0}
	sum, err := v.Add(v2)
	if err != nil {
		t.Fatal(err)
	}
	var _ Vector = sum // Should compile - sum is Vector

	// Test functions work with Vector
	dot, err := DotProduct(v, v2)
	if err != nil {
		t.Fatal(err)
	}
	var _ float64 = dot // Should compile - dot is float64

	// Test that Vector and Vec[float64] are identical
	var v3 Vec[float64] = Vector{5.0, 6.0}
	var v4 Vector = Vec[float64]{7.0, 8.0}
	_, err = v3.Add(v4)
	if err != nil {
		t.Fatal(err)
	}
}

// TestTypeSafety verifies that type mixing is prevented at compile time
// Note: These are commented out as they should NOT compile
// func TestTypeSafety(t *testing.T) {
// 	v32 := Vec[float32]{1.0, 2.0}
// 	v64 := Vec[float64]{3.0, 4.0}
//
// 	// This should NOT compile:
// 	// _, _ = v32.Add(v64)  // Cannot mix float32 and float64
// 	// _, _ = DotProduct(v32, v64)  // Cannot mix types
// }
