package main

import (
	"context"
	"net"

	"github.com/MSSkowron/GoMicroPriceFetcher/proto"
	"google.golang.org/grpc"
)

func makeGRPCServerAndRun(listenAddr string, svc priceFetcher) error {
	grpcServer := NewGRPCServer(&svc)

	options := []grpc.ServerOption{}
	server := grpc.NewServer(options)

	proto.RegisterPriceFetcherServer(server, grpcServer)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	return server.Serve(ln)
}

type GRPCServer struct {
	svc PriceFetcher
	proto.UnimplementedPriceFetcherServer
}

func NewGRPCServer(s PriceFetcher) *GRPCServer {
	return &GRPCServer{
		svc: s,
	}
}

func (s *GRPCServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
	price, err := s.svc.FetchPrice(ctx, req.Ticker)
	if err != nil {
		return nil, err
	}

	return &proto.PriceResponse{
		Ticker: req.Ticker,
		Price:  float32(price),
	}, nil
}
