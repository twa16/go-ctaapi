package bus

import (
	"os"
	"testing"
)

func newTestClient(t *testing.T) *Client {
	apiKey := os.Getenv("CTA_API_KEY")
	if apiKey == "" {
		t.Skip("CTA_API_KEY environment variable not set, skipping integration test")
	}
	return NewClient(apiKey)
}
