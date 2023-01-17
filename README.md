# GoPriceFetcherMicro

GoPriceFetcherMicro is a simple microservice written in Go language to fetch cryprocurrencies price. Prices are currently hard coded in server.go file. There are currently two available: BTC and ETH.
HTTP Server is running on port 3000 and gRPC Server is running on port 4000.

## Technologies

- Go
- PostgreSQL
- Docker

## Requirements

We need to have Docker installed in order to run the application.

## Installation

`git clone https://github.com/MSSkowron/GoPriceFetcherMicro`

## How to run

```
cd GoPriceFetcherMicro
docker-compose up -d
```

## How to use the API

- **API**
  API Server is running on port 3000.
  In order to fetch a price we need to make a HTTP request specyfing _ticker_ as a query parameter.
  For example: http://localhost:3000/?ticker=BTC
  JSON Response Body:

  ```json
  {
    "ticker":    string,
    "price":     float64
  }
  ```

- **gRPC**
  gRPC Server is running on port 4000.
  In order to fetch a price we need to use the _proto.PriceFetcherClient_ struct from the proto package (file service_grpc.pb.go) using the _FetchPrice_ method.
