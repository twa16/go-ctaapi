package bus

import (
	"context"
	"errors"
	"net/url"
	"strconv"
	"strings"
)

// Prediction for a vehicle's arrival or departure.
type Prediction struct {
	Timestamp          string `json:"tmstmp"`
	Type               string `json:"typ"`
	StopName           string `json:"stpnm"`
	StopID             string `json:"stpid"`
	VehicleID          string `json:"vid"`
	DistToStop         int    `json:"dstp"`
	Route              string `json:"rt"`
	RouteDD            string `json:"rtdd"`
	RouteDirection     string `json:"rtdir"`
	Destination        string `json:"des"`
	PredictionTime     string `json:"prdtm"`
	BlockID            string `json:"tablockid"`
	TripID             string `json:"tatripid"`
	OrigTripNo         string `json:"origtatripno"`
	Delayed            bool   `json:"dly"`
	DynamicAction      int    `json:"dyn"`
	Countdown          string `json:"prdctdn"`
	Zone               string `json:"zone"`
	PassengerLoad      string `json:"psgld"`
	ScheduledTripStart int    `json:"stst"`
	ScheduledTripDate  string `json:"stsd"`
	FlagStop           int    `json:"flagstop"`
}

// PredictionsResponse is the response from a GetPredictions request.
type PredictionsResponse struct {
	Body struct {
		Predictions []Prediction `json:"prd"`
		BusTimeResponse
	} `json:"bustime-response"`
}

// GetPredictionsOptions provides options for the GetPredictions method.
type GetPredictionsOptions struct {
	StopIDs      []string
	Routes       []string
	VehicleIDs   []string
	Top          int
	TimestampRes string
	UnixTime     bool
}

// GetPredictions retrieves predictions for stops or vehicles.
func (c *Client) GetPredictions(ctx context.Context, opts *GetPredictionsOptions) (*PredictionsResponse, error) {
	if opts == nil {
		return nil, errors.New("options must be provided")
	}

	isStopIDs := len(opts.StopIDs) > 0
	isVehicleIDs := len(opts.VehicleIDs) > 0

	if (isStopIDs && isVehicleIDs) || (!isStopIDs && !isVehicleIDs) {
		return nil, errors.New("either StopIDs or VehicleIDs must be provided, but not both")
	}

	params := url.Values{}
	if isStopIDs {
		params.Set("stpid", strings.Join(opts.StopIDs, ","))
		if len(opts.Routes) > 0 {
			params.Set("rt", strings.Join(opts.Routes, ","))
		}
	} else {
		params.Set("vid", strings.Join(opts.VehicleIDs, ","))
	}

	if opts.Top > 0 {
		params.Set("top", strconv.Itoa(opts.Top))
	}
	if opts.TimestampRes != "" {
		params.Set("tmres", opts.TimestampRes)
	}
	if opts.UnixTime {
		params.Set("unixTime", "true")
	}

	req, err := c.newRequest("getpredictions", params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	var response PredictionsResponse
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Body.Errors) > 0 {
		return nil, response.Body.Errors[0]
	}

	return &response, nil
}
