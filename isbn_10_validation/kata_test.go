package kata

import (
	"testing"
)

func TestValidISBN10(t *testing.T) {
	cases := []struct {
		isbn string
		want    bool
	}{
		{"1234554321", true},
		{"1293000000", true},
		{"048665088X", true},
		{"1112223339", true},
	}

	for _, c := range cases {
		got := ValidISBN10(c.isbn)
		if got != c.want {
			t.Errorf("ValidISBN10(%s) => %v, want %v", c.isbn, got, c.want)
		}
	}
}
