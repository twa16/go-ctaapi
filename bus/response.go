package bus

import "fmt"

// APIError represents an error returned by the CTA API.
type APIError struct {
	Message string `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("CTA API Error: %s", e.Message)
}

type BusTimeResponse struct {
	Errors []APIError `json:"error"`
}
