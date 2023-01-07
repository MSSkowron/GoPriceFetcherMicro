package main

import (
	"context"
	"fmt"
	"time"
)

var prices = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
}

// PriceFetcher is an interface that can fetch a price.
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher implements the PriceFetcher interface.
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	time.Sleep(100 * time.Millisecond)

	price, ok := prices[ticker]
	if !ok {
		return -1, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}

	return price, nil
}
