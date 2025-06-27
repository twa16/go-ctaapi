package bus

import (
	"context"
	"testing"
)

func TestGetPatterns(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	t.Run("by route", func(t *testing.T) {
		opts := &GetPatternsOptions{
			Route: "20",
		}
		resp, err := client.GetPatterns(ctx, opts)
		if err != nil {
			t.Fatal(err)
		}
		if len(resp.Body.Patterns) == 0 {
			t.Error("expected patterns, got none")
		}
	})

	t.Run("by pattern ids", func(t *testing.T) {
		// From the docs, pattern 954 is for route 20
		pid := "954"
		opts := &GetPatternsOptions{
			PatternIDs: []string{pid},
		}
		resp, err := client.GetPatterns(ctx, opts)
		if err != nil {
			t.Fatal(err)
		}
		if len(resp.Body.Patterns) != 1 {
			t.Errorf("expected 1 pattern, got %d", len(resp.Body.Patterns))
		}
	})

	t.Run("invalid options", func(t *testing.T) {
		_, err := client.GetPatterns(ctx, &GetPatternsOptions{})
		if err == nil {
			t.Error("expected error for empty options")
		}

		_, err = client.GetPatterns(ctx, &GetPatternsOptions{Route: "20", PatternIDs: []string{"954"}})
		if err == nil {
			t.Error("expected error for mutually exclusive options")
		}
	})
}
