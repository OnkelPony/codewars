package kata

import (
	"testing"
)

func TestLastDigit(t *testing.T) {
	tests := []struct {
		Base           string
		Exponent       string
		ExpectedOutput int
	}{
		{"4", "12", 6},
        {"125", "12", 5},
        {"4", "2", 6},
        {"7", "5", 7},
		{"4", "2", 6},
        {"4", "0", 1},
        {"10", "123", 0},
		{"1606938044258990275541962092341162602522202993782792835301376", "2037035976334486086268445688409378161051468393665936250636140449354381299763336706183397376", 6},
	}

	// Loop through each struct in the tests slice
	for _, tt := range tests {
		actualOutput := LastDigit(tt.Base, tt.Exponent)

		// Use the built-in strings package Compare method to compare the expected and actual outputs
		if actualOutput != tt.ExpectedOutput {
			t.Errorf("LastDigit(%s, %s): expected %d but got %d", tt.Base, tt.Exponent, tt.ExpectedOutput, actualOutput)
		}
	}
}
