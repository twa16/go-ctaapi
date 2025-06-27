package bus

import (
	"context"
	"net/url"
	"strconv"
)

type TimeResponse struct {
	Body struct {
		Time string `json:"tm"`
	} `json:"bustime-response"`
}

// GetTimeOptions provides options for the GetTime method
type GetTimeOptions struct {
	// If true, returns the number of milliseconds that have elapsed since
	// 00:00:00 Coordinated Universal Time (UTC), Thursday, 1 January 1970
	UnixTime bool
}

func (c *Client) GetTime(ctx context.Context, opts *GetTimeOptions) (*TimeResponse, error) {
	params := url.Values{}
	if opts != nil {
		if opts.UnixTime {
			params.Set("unixTime", strconv.FormatBool(opts.UnixTime))
		}
	}

	req, err := c.newRequest("gettime", params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	var response TimeResponse
	err = c.do(req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
