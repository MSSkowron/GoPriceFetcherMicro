package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/MSSkowron/GoPriceFetcher/client"
)

func main() {
	client := client.NewClient("http://localhost:3000")
	price, err := client.FetchPrice(context.Background(), "BTCXX")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", price)

	return
	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()

	svc := NewLoggingService(NewMetricsService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}
