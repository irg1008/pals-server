package client

type Client struct {
	URL string
}

func NewClient(url string) *Client {
	return &Client{URL: url}
}
