package kata_test

import (
  "testing"
  "kata"
)

func TestValidBraces(t *testing.T) {
  
  cases := []struct {
    str string
    expected bool
  }{
    {str: "(){}[]", expected: true},
    {str: "([{}])", expected: true},
    {str: "(}", expected: false},
    {str: "[(]", expected: false},
    {str: "[({)](]", expected: false},
    {str: "", expected: true},
    {str: " ", expected: false},
    {str: "abcd", expected: false},
    {str: "(a)", expected: false},
    {str: "(a)b", expected: false},
    {str: "(}", expected: false},
    {str: "([])", expected: true},
    {str: "[({)](]", expected: false},
  }

  for _, tc := range cases {
    t.Run(tc.str, func(t *testing.T) {
      actual := kata.ValidBraces(tc.str)
      if actual != tc.expected {
        t.Errorf("Expected %v, got %v", tc.expected, actual)
      }
    })
  }
}