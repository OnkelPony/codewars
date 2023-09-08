package kata

import (
        "testing"
)

// TestRot13 takes a plaintext input and its expected ciphertext output as parameters
func TestLCS(t *testing.T) {
        // Define a slice of structs containing test inputs and corresponding outputs
        tests := []struct {
                StartIP       string
                EndIP         string
                ExpectedCount int
        }{
                {"10.0.0.0", "10.0.0.50", 50},
                {"20.0.0.10", "20.0.1.0", 246},
        }

        // Loop through each struct in the tests slice
        for _, tt := range tests {
                actualCount := IpsBetween(tt.StartIP, tt.EndIP)

                // Use the built-in strings package Compare method to compare the expected and actual outputs
                if actualCount != tt.ExpectedCount {
                        t.Errorf("IpsBetween(%s, %s): expected %d but got %d", tt.StartIP, tt.EndIP, tt.ExpectedCount, actualCount)
                }
        }
}
