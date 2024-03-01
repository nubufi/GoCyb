package shodan

const (
	API_KEY  = "x59n6O40WKMD2ID2JSIAh7BrKFD0WsPh"
	BASE_URL = "https://api.shodan.io"
)

type Client struct{ ApiKey string }

func New(apiKey string) *Client {
	return &Client{ApiKey: apiKey}
}
