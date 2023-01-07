package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/MSSkowron/GoMicroPriceFetcher/client"
	"github.com/MSSkowron/GoMicroPriceFetcher/proto"
)

func main() {
	var (
		jsonAddr = flag.String("json", ":3000", "listen address the service is running")
		grpcAddr = flag.String("grpc", ":4000", "listen address the service is running")
		svc      = NewLoggingService(NewMetricsService(&priceFetcher{}))
		ctx      = context.Background()
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		time.Sleep(3 * time.Second)

		resp, err := grpcClient.FetchPrice(ctx, &proto.PriceRequest{Ticker: "BTC"})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%+v\n", resp)
	}()

	go func() {
		if err := makeGRPCServerAndRun(*grpcAddr, svc); err != nil {
			log.Fatalln(err)
		}
	}()

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()
}
