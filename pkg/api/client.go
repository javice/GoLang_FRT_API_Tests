// pkg/api/client.go

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	BaseURL = "https://jsonplaceholder.typicode.com"
	Timeout = 10 * time.Second
)

type Client struct {
	BaseURL string
	HTTPClient *http.Client
}

func NewClient() *Client {
	return &Client{
		BaseURL: BaseURL,
		HTTPClient: &http.Client{
			Timeout: Timeout,
		},
	}
}

func (c *Client) SendRequest( method, endpoint string, body interface{}) (*http.Response, error) {

	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)

	var reqBody io.Reader
	if body != nil {
		jsonBody,err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %v", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	return resp, nil
}

