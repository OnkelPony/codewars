package kata

import (
	"reflect"
	"testing"
)

func TestSolequa	(t *testing.T) {
    tests := []struct{
        Input    int
        ExpectedOutput [][]int
    } {
        {5, [][]int{{3, 1}}},
        {12, [][]int{{4, 1}}}, 
		{13, [][]int{{7, 3}}},
		{9005, [][]int{{4503, 2251}, {903, 449}}},
		{9008, [][]int{{1128, 562}}},
		{90002, [][]int{}},
		{1000000, [][]int{{125002, 62499}, {62504, 31248}, {31258, 15621}, {25010, 12495}, {12520, 6240}, {6290, 3105}, {5050, 2475}, {2600, 1200}, {1450, 525}, {1250, 375}, {1000, 0}}},
        // {1000000000, [][]int{}},

    }

    for _, tt := range tests {
        actualOutput := Solequa(tt.Input)

        if !reflect.DeepEqual(tt.ExpectedOutput, actualOutput){
            t.Errorf("Solequa(%d): expected %v but got %#v", tt.Input, tt.ExpectedOutput, actualOutput)
        }
    }
}
