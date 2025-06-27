package bus

import (
	"context"
	"testing"
)

func TestGetRoutes(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	resp, err := client.GetRoutes(ctx)
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Body.Routes) == 0 {
		t.Error("expected routes, got none")
	}
}

func TestGetDirections(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	// Using route "20" as it's a common one
	resp, err := client.GetDirections(ctx, "20")
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.Body.Directions) == 0 {
		t.Error("expected directions for route 20, got none")
	}

	// Test invalid route
	_, err = client.GetDirections(ctx, "INVALID_ROUTE")
	if err == nil {
		t.Error("expected error for invalid route, got nil")
	}
}
