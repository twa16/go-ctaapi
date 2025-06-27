package bus

import (
	"context"
	"errors"
	"net/url"
	"strings"
)

// Vehicle describes a CTA bus.
type Vehicle struct {
	VehicleID          string `json:"vid"`
	Timestamp          string `json:"tmstmp"`
	Latitude           string `json:"lat"`
	Longitude          string `json:"lon"`
	Heading            string `json:"hdg"`
	PatternID          int    `json:"pid"`
	Route              string `json:"rt"`
	Destination        string `json:"des"`
	PatternDist        int    `json:"pdist"`
	Delayed            bool   `json:"dly"`
	Speed              int    `json:"spd"`
	TripID             string `json:"tatripid"`
	BlockID            string `json:"tablockid"`
	OrigTripNo         string `json:"origtatripno"`
	Zone               string `json:"zone"`
	Mode               int    `json:"mode"`
	PassengerLoad      string `json:"psgld"`
	ScheduledTripStart int    `json:"stst"`
	ScheduledTripDate  string `json:"stsd"`
}

// VehiclesResponse is the response from a GetVehicles request.
type VehiclesResponse struct {
	Body struct {
		Vehicles []Vehicle `json:"vehicle"`
		BusTimeResponse
	} `json:"bustime-response"`
}

// GetVehiclesOptions provides options for the GetVehicles method.
type GetVehiclesOptions struct {
	VehicleIDs   []string
	Routes       []string
	TimestampRes string
}

// GetVehicles retrieves vehicle information for all or a subset of vehicles currently being tracked.
func (c *Client) GetVehicles(ctx context.Context, opts *GetVehiclesOptions) (*VehiclesResponse, error) {
	if opts != nil && len(opts.VehicleIDs) > 0 && len(opts.Routes) > 0 {
		return nil, errors.New("VehicleIDs and Routes are mutually exclusive")
	}

	params := url.Values{}
	if opts != nil {
		if len(opts.VehicleIDs) > 0 {
			params.Set("vid", strings.Join(opts.VehicleIDs, ","))
		}
		if len(opts.Routes) > 0 {
			params.Set("rt", strings.Join(opts.Routes, ","))
		}
		if opts.TimestampRes != "" {
			params.Set("tmres", opts.TimestampRes)
		}
	}

	req, err := c.newRequest("getvehicles", params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	var response VehiclesResponse
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Body.Errors) > 0 {
		return nil, response.Body.Errors[0]
	}

	return &response, nil
}
