package main

import (
	"flag"
)

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	flag.Parse()

	svc := NewLoggingService(NewMetricsService(&priceFetcher{}))

	server := NewJSONAPIServer(*listenAddr, svc)
	server.Run()
}

//	client := client.NewClient("http://localhost:3000")
// price, err := client.FetchPrice(context.Background(), "BTC")
// if err != nil {
// 	log.Fatalln(err)
// }

// fmt.Printf("%+v\n", price)

// return
