package adds

import (
	"testing"
)

// TestDetectService verifies that the DetectService function correctly identifies services.
//
// This test checks both the cases where a service is known (exists in the services map) and
// where a service is unknown (does not exist in the services map).
func TestDetectService(t *testing.T) {
	// Define a sample services map for testing
	services := map[int]string{
		80:  "HTTP",
		443: "HTTPS",
		22:  "SSH",
	}

	// Test cases
	tests := []struct {
		port     int
		expected ServiceVersion
	}{
		{80, ServiceVersion{Port: 80, Protocol: "Unknown", Service: "HTTP", Response: "Service Detected"}},
		{443, ServiceVersion{Port: 443, Protocol: "Unknown", Service: "HTTPS", Response: "Service Detected"}},
		{22, ServiceVersion{Port: 22, Protocol: "Unknown", Service: "SSH", Response: "Service Detected"}},
		{9999, ServiceVersion{Port: 9999, Protocol: "Unknown", Service: "Unknown", Response: "Service Not Detected"}},
	}

	// Iterate over test cases and check results
	for _, tt := range tests {
		t.Run(tt.expected.Service, func(t *testing.T) {
			result := DetectService(tt.port, services)
			if result != tt.expected {
				t.Errorf("DetectService(%d) = %v; want %v", tt.port, result, tt.expected)
			}
		})
	}
}
