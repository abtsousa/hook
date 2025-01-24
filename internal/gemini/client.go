package gemini

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client  http.Client
	api_key string
}

const API_PATH = ".api-key"

func NewClient(timeout time.Duration) (Client, err) {
	api_key, err := os.ReadFile(API_PATH)
	if err != nil {
		return Client{}, fmt.Errorf("Please save your API key in the file %s", API_PATH)
	}
	return Client{
		client: http.Client{
			Timeout: timeout,
		},
		api_key: string(api_key),
	}, nil
}
