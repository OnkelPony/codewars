package kata

import (
	"reflect"
	"testing"
)

func TestPickPeaks(t *testing.T) {
	// Define a slice of structs containing test inputs and corresponding outputs
	tests := []struct {
		Input    []int
		Expected PosPeaks
	}{
		{[]int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1}, PosPeaks{Pos: []int{3, 7}, Peaks: []int{6, 3}}},
	}

	// Loop through each struct in the tests slice
	for _, tt := range tests {
		actual := PickPeaks(tt.Input)

		if !reflect.DeepEqual(tt.Expected, actual) {
			t.Errorf("PickPeaks(%v): expected %v but got %v", tt.Input, tt.Expected, actual)
		}
	}
}
