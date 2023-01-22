package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/MSSkowron/GoMicroPriceFetcher/client"
	"github.com/MSSkowron/GoMicroPriceFetcher/proto"
)

func main() {
	var (
		jsonAddr = flag.String("json", ":3000", "listen address the service is running")
		grpcAddr = flag.String("grpc", ":4000", "listen address the service is running")
		svc      = NewLoggingService(NewMetricsService(&priceFetcher{}))
	)
	flag.Parse()

	go func() {
		if err := MakeGRPCServerAndRun(*grpcAddr, svc); err != nil {
			log.Fatalln(err)
		}
	}()

	jsonServer := NewJSONAPIServer(*jsonAddr, svc)
	jsonServer.Run()

	client, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatalln()
	}

	response, err := client.FetchPrice(context.Background(), &proto.PriceRequest{Ticker: "BTC"})
	if err != nil {
		log.Fatalln()
	}

	fmt.Println(response.Price)
}
