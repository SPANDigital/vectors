package vectors

import (
	"reflect"
	"testing"
)

func TestVector_Normalize(t *testing.T) {
	tests := []struct {
		name string
		v    Vector
		want Vector
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.v.Normalize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
