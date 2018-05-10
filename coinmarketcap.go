package coinmarketcap

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://api.coinmarketcap.com/v2/"
)

type Client struct {
	client  *http.Client
	baseURL *url.URL
}

type Metadata struct {
	Timestamp           int     `json:"timestamp"`
	NumCryptocurrencies int     `json:"num_cryptocurrencies"`
	Error               *string `json:"error"`
}

type Response struct {
	Data     interface{} `json:"data"`
	Metadata *Metadata   `json:"metadata"`
}

func NewClient() *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:  http.DefaultClient,
		baseURL: baseURL,
	}

	return c
}

func (c *Client) NewRequest(path string, opt interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	if opt != nil {
		q, err := query.Values(opt)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req := &http.Request{
		Method:     "GET",
		URL:        u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       u.Host,
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := new(Response)
	if v != nil {
		response.Data = v
	}
	err = json.NewDecoder(resp.Body).Decode(response)

	if response.Metadata.Error != nil {
		return response, errors.New(*response.Metadata.Error)
	}

	return response, err
}

func parseID(id interface{}) (string, error) {
	switch v := id.(type) {
	case int:
		return strconv.Itoa(v), nil
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("invalid ID type %#v, the ID must be an int or a string", id)
	}
}
