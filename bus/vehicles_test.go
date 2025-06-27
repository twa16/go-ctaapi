package bus

import (
	"context"
	"testing"
)

func TestGetVehicles(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	t.Run("by route", func(t *testing.T) {
		opts := &GetVehiclesOptions{
			Routes: []string{"20"},
		}
		resp, err := client.GetVehicles(ctx, opts)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("Found %d vehicles on route 20", len(resp.Body.Vehicles))
		if len(resp.Body.Vehicles) > 0 {
			// Test by vehicle ID with a vehicle we just found
			vid := resp.Body.Vehicles[0].VehicleID
			t.Run("by vehicle ids", func(t *testing.T) {
				opts := &GetVehiclesOptions{
					VehicleIDs: []string{vid},
				}
				resp, err := client.GetVehicles(ctx, opts)
				if err != nil {
					t.Fatal(err)
				}
				if len(resp.Body.Vehicles) != 1 {
					t.Errorf("expected 1 vehicle, got %d", len(resp.Body.Vehicles))
				}
			})
		} else {
			t.Log("Skipping vehicle by ID test, no vehicles found on route 20")
		}
	})

	t.Run("invalid options", func(t *testing.T) {
		_, err := client.GetVehicles(ctx, &GetVehiclesOptions{Routes: []string{"20"}, VehicleIDs: []string{"1234"}})
		if err == nil {
			t.Error("expected error for mutually exclusive options")
		}
	})
}
