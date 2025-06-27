package bus

import (
	"context"
	"testing"
)

func TestGetPredictions(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	t.Run("by stop ids", func(t *testing.T) {
		opts := &GetPredictionsOptions{
			StopIDs: []string{"456"},
		}
		resp, err := client.GetPredictions(ctx, opts)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("Found %d predictions for stop 456", len(resp.Body.Predictions))
		if len(resp.Body.Predictions) > 0 {
			vid := resp.Body.Predictions[0].VehicleID
			t.Run("by vehicle id", func(t *testing.T) {
				opts := &GetPredictionsOptions{
					VehicleIDs: []string{vid},
				}
				resp, err := client.GetPredictions(ctx, opts)
				if err != nil {
					t.Fatal(err)
				}
				if len(resp.Body.Predictions) == 0 {
					t.Error("expected predictions, got none")
				}
			})
		} else {
			t.Log("Skipping prediction by vehicle ID test, no predictions found for stop 456")
		}
	})

	t.Run("invalid options", func(t *testing.T) {
		_, err := client.GetPredictions(ctx, &GetPredictionsOptions{})
		if err == nil {
			t.Error("expected error for empty options")
		}

		_, err = client.GetPredictions(ctx, &GetPredictionsOptions{StopIDs: []string{"456"}, VehicleIDs: []string{"1234"}})
		if err == nil {
			t.Error("expected error for mutually exclusive options")
		}
	})
}
