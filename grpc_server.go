package main

import (
	"context"
	"math/rand"
	"net"

	"github.com/MSSkowron/GoMicroPriceFetcher/proto"
	"google.golang.org/grpc"
)

func MakeGRPCServerAndRun(listenAddr string, svc PriceFetcher) error {
	grpcServer := newGRPCServer(svc)

	options := []grpc.ServerOption{}
	server := grpc.NewServer(options...)

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

func newGRPCServer(s PriceFetcher) *GRPCServer {
	return &GRPCServer{
		svc: s,
	}
}

func (s *GRPCServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
	ctx = context.WithValue(ctx, "requestID", rand.Intn(100000))

	price, err := s.svc.FetchPrice(ctx, req.Ticker)
	if err != nil {
		return nil, err
	}

	return &proto.PriceResponse{
		Ticker: req.Ticker,
		Price:  float32(price),
	}, nil
}
