package vectors

import "testing"

func TestVector_IsNormalized(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want bool
	}{
		{
			name: "Normalized",
			v:    Vector{1, 0, 0},
			want: true,
		},
		{
			name: "Not Normalized",
			v:    Vector{1, 1, 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.IsNormalized(); got != tt.want {
				t.Errorf("IsNormalized() = %v, want %v", got, tt.want)
			}
		})
	}
}
