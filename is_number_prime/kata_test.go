package kata

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	tests := []struct {
		n int
		want bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{8, false},
		{9, false},
		{10, false},
	}

	for _, tt := range tests {
		got := IsPrime(tt.n)
		if got != tt.want {
			t.Errorf("IsPrime(%d) = %v, want %v", tt.n, got, tt.want)
		}
	}
}
