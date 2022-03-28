package httpClient

import (
	"io"
	"net/http"
	"time"
)

// Response represents an HTTP response. It contains some of the fields in http.Response.
// It is meant to be used as a simpler structure instead of http.Response.
type Response struct {
	Body          []byte
	ContentLength int64
	Status        string
	StatusCode    int
}

// Client is a wrapper around an http.Client.
type Client struct {
	http *http.Client
}

// New returns a new instance of a Client.
func New() *Client {
	return &Client{
		http: &http.Client{
			// Time out a request if it takes 30 seconds or more
			Timeout: time.Duration(30) * time.Second,
		},
	}
}

// Get issues a GET request to the specified URL.
func (c *Client) Get(url string) (*Response, error) {
	resp, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &Response{
		Body:          respBytes,
		ContentLength: resp.ContentLength,
		Status:        resp.Status,
		StatusCode:    resp.StatusCode,
	}

	return response, nil
}
