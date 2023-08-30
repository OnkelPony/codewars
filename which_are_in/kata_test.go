package kata_test

import (
	"fmt"
	"kata"
	"testing"
)

func TestInArray(t *testing.T) {
	array1 := []string{"hello", "world", "rl", "good"}
	array2 := []string{"hello", "world", "worlds", "carlson"}

	result := kata.InArray(array1, array2)

	expected := []string{"hello", "rl", "world"}

	if fmt.Sprint(result) != fmt.Sprint(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
