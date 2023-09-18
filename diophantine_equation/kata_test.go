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
    }

    for _, tt := range tests {
        actualOutput := Solequa(tt.Input)

        if !reflect.DeepEqual(tt.ExpectedOutput, actualOutput){
            t.Errorf("Solequa(%d): expected %v but got %v", tt.Input, tt.ExpectedOutput, actualOutput)
        }
    }
}
