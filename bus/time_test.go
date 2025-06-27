package bus

import (
	"context"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	t.Run("normal time", func(t *testing.T) {
		resp, err := client.GetTime(ctx, nil)
		if err != nil {
			t.Fatal(err)
		}
		if resp.Body.Time == "" {
			t.Error("expected time, got empty string")
		}
		// Try to parse it
		_, err = time.Parse("20060102 15:04:05", resp.Body.Time)
		if err != nil {
			t.Errorf("failed to parse time: %v", err)
		}
	})

	t.Run("unix time", func(t *testing.T) {
		resp, err := client.GetTime(ctx, &GetTimeOptions{UnixTime: true})
		if err != nil {
			t.Fatal(err)
		}
		if resp.Body.Time == "" {
			t.Error("expected unix time, got empty string")
		}
	})
}
