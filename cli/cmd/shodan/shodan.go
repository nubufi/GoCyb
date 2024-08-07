package shodan

const (
	API_KEY  = ""
	BASE_URL = "https://api.shodan.io"
)

type Client struct{ ApiKey string }

func New(apiKey string) *Client {
	return &Client{ApiKey: apiKey}
}
