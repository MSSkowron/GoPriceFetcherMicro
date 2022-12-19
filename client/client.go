package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MSSkowron/GoPriceFetcher/types"
)

type Client struct {
	endpoint string
}

func NewClient(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s?ticker=%s", c.endpoint, ticker)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("service responded with non OK status code: %d", res.StatusCode)
	}

	priceResponse := &types.PriceResponse{}
	if err := json.NewDecoder(res.Body).Decode(priceResponse); err != nil {
		return nil, err
	}

	return priceResponse, nil
}
