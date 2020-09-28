package client

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

// SpreedlyAPIURL is a const
const SpreedlyAPIURL = "https://core.spreedly.com/v1/"

// Client is the client API used to communicate with Spreedly
type Client struct {
	BaseURL *url.URL
	Auth    Auth

	httpClient *http.Client
}

// Auth ..
type Auth struct {
	AccessKey    string
	AccessSecret string
}

// WriteRequest ...Thinking about separating Read & Write API calls
func (c *Client) WriteRequest(method string, path string, data []byte) (*http.Response, error) {
	var body io.Reader

	body = bytes.NewReader(data)

	url, err := url.Parse(SpreedlyAPIURL + path)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		fmt.Println(err)
	}

	req.SetBasicAuth(c.Auth.AccessKey, c.Auth.AccessSecret)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("User-Agent", "Spreedly-CLI")

	if c.httpClient == nil {
		c.httpClient = newHTTPClient()
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ReadRequest ...
func (c *Client) ReadRequest(method string, path string) (*http.Response, error) {
	var body io.Reader

	url, err := url.Parse(SpreedlyAPIURL + path)
	fmt.Println(url.String())
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		fmt.Println(err)
	}
	req.SetBasicAuth(c.Auth.AccessKey, c.Auth.AccessSecret)
	req.Header.Set("content-type", "application/json")
	req.Header.Set("User-Agent", "Spreedly-CLI")

	if c.httpClient == nil {
		c.httpClient = newHTTPClient()
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func newHTTPClient() *http.Client {
	var httpTransport *http.Transport

	httpTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
	}

	return &http.Client{
		Transport: httpTransport,
	}
}
