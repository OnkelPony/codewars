package kata
import (
    "strings"
    "testing"
)

// TestRot13 takes a plaintext input and its expected ciphertext output as parameters
func TestRot13(t *testing.T) {
    // Define a slice of structs containing test inputs and corresponding outputs
    tests := []struct{
        Input    string
        ExpectedOutput string
    } {
        {"Hello World!", "Uryyb Jbeyq!"} ,
        {"My name is John", "Zl anzr vf Wbua"}, 
    }

    // Loop through each struct in the tests slice
    for _, tt := range tests {
        actualOutput := Rot13(tt.Input)

        // Use the built-in strings package Compare method to compare the expected and actual outputs
        if !strings.EqualFold(actualOutput, tt.ExpectedOutput){
            t.Errorf("Rot13(%s): expected %s but got %s", tt.Input, tt.ExpectedOutput, actualOutput)
        }
    }
}