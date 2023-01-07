package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MSSkowron/GoMicroPriceFetcher/proto"
	"github.com/MSSkowron/GoMicroPriceFetcher/types"
	"google.golang.org/grpc"
)

func NewGRPCClient(remoteAddr string) (proto.PriceFetcherClient, error) {
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := proto.NewPriceFetcherClient(conn)
	return c, nil
}

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
		httpErr := map[string]any{}
		if err := json.NewDecoder(res.Body).Decode(&httpErr); err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("service responded with non OK status code: %d - %s", res.StatusCode, httpErr)
	}

	priceResponse := &types.PriceResponse{}
	if err := json.NewDecoder(res.Body).Decode(priceResponse); err != nil {
		return nil, err
	}

	return priceResponse, nil
}
