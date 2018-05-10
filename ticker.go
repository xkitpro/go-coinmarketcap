package coinmarketcap

import (
	"context"
	"fmt"
)

type Ticker struct {
	ID                int                    `json:"id"`
	Name              string                 `json:"name"`
	Symbol            string                 `json:"symbol"`
	WebsiteSlug       string                 `json:"website_slug"`
	Rank              int                    `json:"rank"`
	CirculatingSupply float32                `json:"circulating_supply"`
	TotalSupply       float32                `json:"total_supply"`
	MaxSupply         float32                `json:"max_supply"`
	Quotes            map[string]TickerQuote `json:"quotes"`
	LastUpdated       int                    `json:"last_updated"`
}

type TickerQuote struct {
	Price float64 `json:"price"`
}

type TickerList map[string]*Ticker

type ListTickerOptions struct {
	Limit   int    `url:"limit,omitempty"`
	Start   int    `url:"start,omitempty"`
	Convert string `url:"convert,omitempty"`
}

func (c *Client) ListTicker(ctx context.Context, opt *ListTickerOptions) (TickerList, *Response, error) {
	req, err := c.NewRequest("ticker/", opt)
	if err != nil {
		return nil, nil, err
	}

	var tickers TickerList

	resp, err := c.Do(ctx, req, &tickers)
	if err != nil {
		return nil, resp, err
	}

	return tickers, resp, nil
}

type GetTickerOptions struct {
	Convert string `url:"convert,omitempty"`
}

func (c *Client) GetTicker(ctx context.Context, id interface{}, opt *GetTickerOptions) (*Ticker, *Response, error) {
	tickerID, err := parseID(id)
	if err != nil {
		return nil, nil, err
	}

	req, err := c.NewRequest(fmt.Sprintf("ticker/%s", tickerID), opt)
	if err != nil {
		return nil, nil, err
	}

	ticker := new(Ticker)

	resp, err := c.Do(ctx, req, ticker)

	return ticker, resp, err
}
