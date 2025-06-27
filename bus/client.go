package bus

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://www.ctabustracker.com/bustime/api/v3/"
)

type Client struct {
	apiKey  string
	BaseURL *url.URL
	client  *http.Client
}

func NewClient(apiKey string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	return &Client{
		apiKey:  apiKey,
		BaseURL: baseURL,
		client:  http.DefaultClient,
	}
}

func (c *Client) newRequest(endpoint string, queryParams url.Values) (*http.Request, error) {
	u, err := c.BaseURL.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("key", c.apiKey)
	q.Set("format", "json")

	for key, values := range queryParams {
		for _, value := range values {
			q.Add(key, value)
		}
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}
