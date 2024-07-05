package vectors

import (
	"math"
	"testing"
)

func TestCosineSimilarity(t *testing.T) {
	tests := []struct {
		name      string
		a         Vector
		b         Vector
		want      float64
		expectErr bool
	}{
		{
			name:      "identical vectors",
			a:         Vector{1, 2, 3},
			b:         Vector{1, 2, 3},
			want:      1,
			expectErr: false,
		},
		{
			name:      "orthogonal vectors",
			a:         Vector{1, 0, 0},
			b:         Vector{0, 1, 0},
			want:      0,
			expectErr: false,
		},
		{
			name:      "different length vectors",
			a:         Vector{1, 2, 3},
			b:         Vector{1, 2},
			want:      0.5976143046671968,
			expectErr: false,
		},
		{
			name:      "zero vector a",
			a:         Vector{0, 0, 0},
			b:         Vector{1, 2, 3},
			want:      0,
			expectErr: true,
		},
		{
			name:      "zero vector b",
			a:         Vector{1, 2, 3},
			b:         Vector{0, 0, 0},
			want:      0,
			expectErr: true,
		},
		{
			name:      "both zero vectors",
			a:         Vector{0, 0, 0},
			b:         Vector{0, 0, 0},
			want:      0,
			expectErr: true,
		},
		{
			name:      "empty vectors",
			a:         Vector{},
			b:         Vector{},
			want:      0,
			expectErr: true,
		},
		{
			name:      "single element vectors",
			a:         Vector{1},
			b:         Vector{1},
			want:      1,
			expectErr: false,
		},
		{
			name:      "single element vectors with zero",
			a:         Vector{0},
			b:         Vector{1},
			want:      0,
			expectErr: true,
		},
		{
			name:      "longer vectors",
			a:         Vector{1, 2, 3, 4, 5},
			b:         Vector{5, 4, 3, 2, 1},
			want:      0.6363636363636364,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CosineSimilarity(tt.a, tt.b)
			if (err != nil) != tt.expectErr {
				t.Errorf("CosineSimilarity() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !tt.expectErr && math.Abs(got-tt.want) > 1e-9 {
				t.Errorf("CosineSimilarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
