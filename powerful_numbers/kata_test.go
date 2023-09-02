package kata_test

import (
	"kata"
	"reflect"
	"testing"
)



func TestGetNextDigPow(t *testing.T) {
	expected := []uint64{1, 4, 64}
	actual := kata.GetNextDigPow(uint64(123), []uint64{1, 2, 27})
	if !reflect.DeepEqual(expected, actual)  {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}