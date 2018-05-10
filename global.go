package coinmarketcap

import "context"

type Global struct {
	ActiveCryptocurrencies       int                    `json:"active_cryptocurrencies"`
	ActiveMarkets                int                    `json:"active_markets"`
	BitcoinPercentageOfMarketCap float32                `json:"bitcoin_percentage_of_market_cap"`
	Quotes                       map[string]GlobalQuote `json:"quotes"`
	LastUpdated                  int                    `json:"last_updated"`
}

type GlobalQuote struct {
	TotalMarketCap float32 `json:"total_market_cap"`
	TotalVolume24h float32 `json:"total_volume_24h"`
}

type GetGlobalOptions struct {
	Convert string `url:"convert,omitempty"`
}

func (c *Client) GetGlobal(ctx context.Context, opt *GetGlobalOptions) (*Global, *Response, error) {
	req, err := c.NewRequest("global/", opt)
	if err != nil {
		return nil, nil, err
	}

	global := new(Global)

	resp, err := c.Do(ctx, req, global)
	if err != nil {
		return nil, resp, err
	}

	return global, resp, nil
}
