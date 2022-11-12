package geocodio

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

var apiBaseUrl = &url.URL{
	Scheme: "https", Host: "api.geocod.io", Path: "/v1.7/",
}

type Client struct {
	BaseUrl *url.URL
	Client  *http.Client
	Key     string
}

func New(key string) *Client {
	return &Client{Key: key}
}

func (c *Client) baseURL() *url.URL {
	if c.BaseUrl == nil {
		return apiBaseUrl
	}

	return c.BaseUrl

}

func (c *Client) client() *http.Client {
	client := c.Client
	if client == nil {
		client = http.DefaultClient
	}
	return client
}

func (c *Client) post(path string, payload, data interface{}) error {
	return c.request(http.MethodPost, path, payload, data)
}

func (c *Client) request(method, path string, payload, data interface{}) error {
	if c.Key == "" {
		return errors.New("missing API Key")
	}

	v := url.Values{}
	v.Set("api_key", c.Key)

	request := &http.Request{
		Method: method,
		Header: http.Header{},
		URL:    c.baseURL().ResolveReference(&url.URL{Path: path, RawQuery: v.Encode()}),
	}

	if payload != nil {
		body, err := json.Marshal(payload)
		if err != nil {
			return err
		}

		request.Header.Add("Content-Type", "application/json")
		request.Body = io.NopCloser(bytes.NewReader(body))
	}

	response, err := c.client().Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(data)

}
