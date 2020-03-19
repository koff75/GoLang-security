package shodan

//	TIP: avoid this kind of function call:
//	     func func(token, url string) { --snip-- }
// Methods on the Client struct, which allows the interrogation on the instance
// Call: func (s *Client) APIInfo() { --snip-- }
// API key through s.apiKey, same for BaseURL

const BaseURL = "https://api.shodan.io"

type Client struct {
	apiKey string
}

// Helper function taking the API token as input
// and returning an initialized Client instance
func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
