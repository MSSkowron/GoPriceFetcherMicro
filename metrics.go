package main

import (
	"context"
	"fmt"
)

type MetricsService struct {
	next PriceFetcher
}

func NewMetricsService(n PriceFetcher) *MetricsService {
	return &MetricsService{
		next: n,
	}
}

func (s *MetricsService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	fmt.Println("doing metrics stuff here")
	return s.next.FetchPrice(ctx, ticker)
}
