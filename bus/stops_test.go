package bus

import (
	"context"
	"testing"
)

func TestGetStops(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	var stopIDs []string
	t.Run("by route and direction", func(t *testing.T) {
		opts := &GetStopsOptions{
			Route:     "20",
			Direction: "Eastbound",
		}
		resp, err := client.GetStops(ctx, opts)
		if err != nil {
			t.Fatal(err)
		}
		if len(resp.Body.Stops) == 0 {
			t.Fatal("expected stops, got none")
		}
		// Save some stop IDs for the next test
		if len(resp.Body.Stops) > 0 {
			stopIDs = append(stopIDs, resp.Body.Stops[0].ID)
		}
		if len(resp.Body.Stops) > 1 {
			stopIDs = append(stopIDs, resp.Body.Stops[1].ID)
		}
	})

	if len(stopIDs) > 0 {
		t.Run("by stop ids", func(t *testing.T) {
			opts := &GetStopsOptions{
				StopIDs: stopIDs,
			}
			resp, err := client.GetStops(ctx, opts)
			if err != nil {
				t.Fatal(err)
			}
			if len(resp.Body.Stops) != len(stopIDs) {
				t.Errorf("expected %d stops, got %d", len(stopIDs), len(resp.Body.Stops))
			}
		})
	} else {
		t.Log("Skipping stops by ID test, no stops found for route 20 Eastbound")
	}

	t.Run("invalid options", func(t *testing.T) {
		_, err := client.GetStops(ctx, &GetStopsOptions{})
		if err == nil {
			t.Error("expected error for empty options")
		}

		_, err = client.GetStops(ctx, &GetStopsOptions{Route: "20", Direction: "Eastbound", StopIDs: []string{"456"}})
		if err == nil {
			t.Error("expected error for mutually exclusive options")
		}
	})
}
