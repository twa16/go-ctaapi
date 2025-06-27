package bus

import (
	"context"
	"errors"
	"net/url"
	"strings"
)

// Point is a geo-positional point in a pattern.
type Point struct {
	Sequence  int     `json:"seq"`
	Type      string  `json:"typ"`
	StopID    string  `json:"stpid"`
	StopName  string  `json:"stpnm"`
	Dist      float64 `json:"pdist"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

// Pattern is a sequence of points that form a path for a vehicle.
type Pattern struct {
	ID           int     `json:"pid"`
	Length       float64 `json:"ln"`
	Direction    string  `json:"rtdir"`
	Points       []Point `json:"pt"`
	DetourID     string  `json:"dtrid"`
	DetourPoints []Point `json:"dtrpt"`
}

// PatternsResponse is the response from a GetPatterns request.
type PatternsResponse struct {
	Body struct {
		Patterns []Pattern `json:"ptr"`
		BusTimeResponse
	} `json:"bustime-response"`
}

// GetPatternsOptions provides options for the GetPatterns method.
type GetPatternsOptions struct {
	PatternIDs []string
	Route      string
}

// GetPatterns retrieves the set of geo-positional points that define a pattern.
func (c *Client) GetPatterns(ctx context.Context, opts *GetPatternsOptions) (*PatternsResponse, error) {
	if opts == nil {
		return nil, errors.New("options must be provided")
	}

	isPIDs := len(opts.PatternIDs) > 0
	isRoute := opts.Route != ""

	if (isPIDs && isRoute) || (!isPIDs && !isRoute) {
		return nil, errors.New("either PatternIDs or Route must be provided, but not both")
	}

	params := url.Values{}
	if isPIDs {
		params.Set("pid", strings.Join(opts.PatternIDs, ","))
	} else {
		params.Set("rt", opts.Route)
	}

	req, err := c.newRequest("getpatterns", params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	var response PatternsResponse
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Body.Errors) > 0 {
		return nil, response.Body.Errors[0]
	}

	return &response, nil
}
