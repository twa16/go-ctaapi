package train

import (
	"net/http"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

func (c *Client) AddAuthKey(req *http.Request) {
	vals := req.URL.Query()
	vals.Add("key", c.apiKey)
	req.URL.RawQuery = vals.Encode()
}