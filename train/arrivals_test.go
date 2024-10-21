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

func TestArrivals(t *testing.T) {
	setup()
	arrivalReq := ArrivalsRequest{
		MapID: "40530",
	}
	resp, err := conn.GetArrivalsAtStation(arrivalReq)
	if err != nil {
		t.Error(err)
	}

	for _, arrival := range resp.Eta {
		fmt.Printf("%s Train heading to %s is arriving at %s\n",
			expandTrainRouteName(arrival.AbbrRouteName),
			arrival.DestinationName,
			arrival.ArrivalTimeParsed)
	}
}
