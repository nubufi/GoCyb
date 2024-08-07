package adds

const (
	Open     = "Open"
	Closed   = "Closed"
	Filtered = "Filtered"
)

type ServiceVersion struct {
	Port     int
	Protocol string
	Service  string
	Response string
}

func DetectService(port int, services map[int]string) ServiceVersion {
	if svc, ok := services[port]; ok {
		return ServiceVersion{Port: port, Protocol: "Unknown", Service: svc, Response: "Service Detected"}
	}
	return ServiceVersion{Port: port, Protocol: "Unknown", Service: "Unknown", Response: "Service Not Detected"}
}
