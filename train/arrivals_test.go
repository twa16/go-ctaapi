package train

import (
	"fmt"
	"os"
	"testing"
)

var conn *Client

func setup() {
	conn = NewClient(os.Getenv("API_KEY"))
}

func TestGetStopsFromChicagoData(t *testing.T) {
	stops, err := GetStopsFromChicagoData()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Got %d stops\n", len(stops))

	filteredStops := GetStopsByName(stops, "Clark/Division")
	for _, stop := range filteredStops {
		fmt.Printf("Got stop %s: %s\n", stop.StopID, stop.StopName)
	}
}

func TestArrivals(t *testing.T) {
	setup()
	arrivalReq := ArrivalsRequest{
		MapID: "40630",
	}
	resp, err := conn.GetArrivalsAtStation(arrivalReq)
	if err != nil {
		t.Fatal(err)
	}

	for _, arrival := range resp.Eta {
		fmt.Printf("%s Train heading to %s is arriving at %s\n",
			expandTrainRouteName(arrival.AbbrRouteName),
			arrival.DestinationName,
			arrival.ArrivalTimeParsed)
	}
}
