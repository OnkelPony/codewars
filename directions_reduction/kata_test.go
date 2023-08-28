package kata_test

import (
    "reflect"
    "testing"

    "kata"
)

func TestDirReduc(t *testing.T) {
    tests := []struct {
        name     string
        input    []string
        expected []string
    }{
        {
            name:  "no duplicates",
            input: []string{"NORTH", "SOUTH", "WEST", "EAST"},
            expected: []string{},
        },
        {
            name:  "duplicates in middle",
            input: []string{"NORTH", "SOUTH", "NORTH", "WEST", "EAST"},
            expected: []string{"NORTH"},
        },
        {
            name:  "duplicates at beginning",
            input: []string{"NORTH", "NORTH", "SOUTH", "WEST", "EAST"},
            expected: []string{"NORTH"},
        },
        {
            name:  "duplicates at end",
            input: []string{"NORTH", "SOUTH", "WEST", "EAST", "EAST"},
            expected: []string{"EAST"},
        },
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            actual := kata.DirReduc(tc.input)
            if !reflect.DeepEqual(actual, tc.expected) {
                t.Errorf("Expected %v, got %v", tc.expected, actual)
            }
        })
    }
}
