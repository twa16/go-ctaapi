package bus

import (
	"context"
	"errors"
	"net/url"
	"strings"
)

// Stop describes a CTA bus stop.
type Stop struct {
	ID        string   `json:"stpid"`
	Name      string   `json:"stpnm"`
	Latitude  float64  `json:"lat"`
	Longitude float64  `json:"lon"`
	DetourAdd []string `json:"dtradd"`
	DetourRem []string `json:"dtrrem"`
	Ada       bool     `json:"ada"`
}

// StopsResponse is the response from a GetStops request.
type StopsResponse struct {
	Body struct {
		Stops []Stop `json:"stops"`
		BusTimeResponse
	} `json:"bustime-response"`
}

// GetStopsOptions provides options for the GetStops method.
type GetStopsOptions struct {
	Route     string
	Direction string
	StopIDs   []string
}

// GetStops retrieves the set of stops for the specified route and direction, or for a list of stop IDs.
func (c *Client) GetStops(ctx context.Context, opts *GetStopsOptions) (*StopsResponse, error) {
	if opts == nil {
		return nil, errors.New("options must be provided")
	}

	isRouteDir := opts.Route != "" && opts.Direction != ""
	isStopIDs := len(opts.StopIDs) > 0

	if (isRouteDir && isStopIDs) || (!isRouteDir && !isStopIDs) {
		return nil, errors.New("either Route and Direction must be provided, or StopIDs must be provided, but not both")
	}

	params := url.Values{}
	if isRouteDir {
		params.Set("rt", opts.Route)
		params.Set("dir", opts.Direction)
	} else {
		params.Set("stpid", strings.Join(opts.StopIDs, ","))
	}

	req, err := c.newRequest("getstops", params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	var response StopsResponse
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Body.Errors) > 0 {
		return nil, response.Body.Errors[0]
	}

	return &response, nil
}
