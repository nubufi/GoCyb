package adds

// Constants representing the state of a port.
//
// Open: The port is open and accepting connections.
// Closed: The port is closed and not accepting connections.
// Filtered: The port is filtered, and it's unclear whether it is open or closed.
const (
	Open     = "Open"
	Closed   = "Closed"
	Filtered = "Filtered"
)

// ServiceVersion holds information about a detected service.
//
// Fields:
// - Port: The port number where the service is running.
// - Protocol: The protocol used by the service (e.g., "TCP", "UDP").
// - Service: The name of the service detected (e.g., "HTTP", "SSH").
// - Response: A message indicating whether the service was detected.
type ServiceVersion struct {
	Port     int
	Protocol string
	Service  string
	Response string
}

// DetectService identifies the service running on a given port.
//
// Parameters:
// - port: The port number to check for a service.
// - services: A map of known services with port numbers as keys and service names as values.
//
// Returns:
// - ServiceVersion: A struct containing details about the detected service.
//
// The function checks if the given port number exists in the provided services map.
// If found, it returns a ServiceVersion struct with the service name and a "Service Detected" response.
// If not found, it returns a ServiceVersion struct with the service name as "Unknown" and a "Service Not Detected" response.
func DetectService(port int, services map[int]string) ServiceVersion {
	if svc, ok := services[port]; ok {
		return ServiceVersion{Port: port, Protocol: "Unknown", Service: svc, Response: "Service Detected"}
	}
	return ServiceVersion{Port: port, Protocol: "Unknown", Service: "Unknown", Response: "Service Not Detected"}
}
