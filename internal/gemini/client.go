package gemini

import (
	"net/http"
	"time"
)

// A simple struct to hold client data.
type Client struct {
	client http.Client
}

// Creates a new HTTP client to use for the API.
func NewClient(timeout time.Duration) (Client, error) {
	return Client{
		client: http.Client{
			Timeout: timeout,
		},
	}, nil
}
