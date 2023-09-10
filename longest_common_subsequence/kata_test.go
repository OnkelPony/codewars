package kata

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestLCS(t *testing.T) {
    tests := []struct {
        s1, s2, expected string
    }{
        {"ABCD", "ACDE", "ACD"},
        {"hello", "world", "l"},
        {"pradedecek", "prasklina", "prak"},
        {"132535365", "123456789", "12356"},
    }

    for _, test := range tests {
        output := LCS(test.s1, test.s2)
        assert.Equal(t, test.expected, output, "LCS(%s, %s) should be %s", test.s1, test.s2, test.expected)
    }
}