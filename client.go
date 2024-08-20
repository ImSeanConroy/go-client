package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/tidwall/gjson"
)

// Client is a simple HTTP client for making requests.
type Client struct {
	BaseURL    string
	HttpClient *http.Client
	Token      string // Bearer token for authentication
}

// NewClient creates a new HTTP client instance.
func NewClient(baseURL, token string) (*Client, error) {
	client := &Client{
		BaseURL:    baseURL,
		HttpClient: &http.Client{},
		Token:      token,
	}
	return client, nil
}

// Do performs an HTTP request and returns the response as a gjson.Result.
func (c *Client) Do(method, path string, body interface{}, headers map[string]string) (gjson.Result, error) {
	url := c.BaseURL + path

	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return gjson.Result{}, err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return gjson.Result{}, err
	}

	// Add Bearer token to headers
	if c.Token != "" {
		req.Header.Add("Authorization", "Bearer "+c.Token)
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return gjson.Result{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return gjson.Result{}, errors.New("request failed with status: " + resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return gjson.Result{}, err
	}

	return gjson.ParseBytes(bodyBytes), nil
}

// Post performs a POST request and returns the response as a gjson.Result.
func (c *Client) Post(path string, body interface{}) (gjson.Result, error) {
	return c.Do(http.MethodPost, path, body, map[string]string{"Content-Type": "application/json"})
}

// Get performs a GET request and returns the response as a gjson.Result.
func (c *Client) Get(path string) (gjson.Result, error) {
	return c.Do(http.MethodGet, path, nil, nil)
}

// Put performs a PUT request and returns the response as a gjson.Result.
func (c *Client) Put(path string, body interface{}) (gjson.Result, error) {
	return c.Do(http.MethodPut, path, body, map[string]string{"Content-Type": "application/json"})
}

// Patch performs a PATCH request and returns the response as a gjson.Result.
func (c *Client) Patch(path string, body interface{}) (gjson.Result, error) {
	return c.Do(http.MethodPatch, path, body, map[string]string{"Content-Type": "application/json"})
}

// Delete performs a DELETE request and returns the response as a gjson.Result.
func (c *Client) Delete(path string) (gjson.Result, error) {
	return c.Do(http.MethodDelete, path, nil, nil)
}
