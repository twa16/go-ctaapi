package bus

import (
	"context"
	"net/url"
)

// Route describes a CTA bus route.
type Route struct {
	ID      string `json:"rt"`
	Name    string `json:"rtnm"`
	Color   string `json:"rtclr"`
	RouteDD string `json:"rtdd"`
}

// RoutesResponse is the response from a GetRoutes request.
type RoutesResponse struct {
	Body struct {
		Routes []Route `json:"routes"`
		BusTimeResponse
	} `json:"bustime-response"`
}

// GetRoutes retrieves the set of routes serviced by the system.
func (c *Client) GetRoutes(ctx context.Context) (*RoutesResponse, error) {
	req, err := c.newRequest("getroutes", nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	var response RoutesResponse
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Body.Errors) > 0 {
		return nil, response.Body.Errors[0]
	}

	return &response, nil
}

// Direction describes a direction of travel for a route.
type Direction struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// DirectionsResponse is the response from a GetDirections request.
type DirectionsResponse struct {
	Body struct {
		Directions []Direction `json:"directions"`
		BusTimeResponse
	} `json:"bustime-response"`
}

// GetDirections retrieves the set of directions for the specified route.
func (c *Client) GetDirections(ctx context.Context, route string) (*DirectionsResponse, error) {
	params := url.Values{}
	params.Set("rt", route)

	req, err := c.newRequest("getdirections", params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	var response DirectionsResponse
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	if len(response.Body.Errors) > 0 {
		return nil, response.Body.Errors[0]
	}

	return &response, nil
}
