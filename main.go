package main

import (
	"flag"
	"log"
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

	NewRESTServer(*jsonAddr, svc).Run()
}
