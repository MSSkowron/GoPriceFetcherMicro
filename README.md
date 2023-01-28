# GoPriceFetcherMicro

GoPriceFetcherMicro is a simple application written in Go language to fetch cryprocurrencies price. Prices are currently hard coded in server.go file. There are currently two available: BTC and ETH.
REST Server is running on port 3000 while gRPC server is running on port 4000.

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

## How to use it?

- *REST Client* \
We can use _Client_ struct from the _client_ package. \
_NewClient_ creates a new instance of the _Client_ struct. We only need to pass the server listening address. In this case it is ":3000". \
In order to fetch a price use _FetchPrice_ method on the client instance. \
Of course in order to fetch a price we need to make a HTTP request specyfing _ticker_ as a request's query parameter. \
For example: http://localhost:3000/?ticker=BTC \
JSON Response Body:
  ```json
  {
    "ticker":    "string",
    "price":     "float64"
  }
  ```

- *gRPC Client* \
We can use _proto.PriceFetcherClient_ struct from the _client_ package. \
_NewGRPCClient_ creates a new instance of the _proto.PriceFetcherClient_ struct. We only need to pass the server listening address. In this case it is ":4000". \
In order to fetch a price use _FetchPrice_ method on the _proto.PriceFetcherClient_ instance.
