package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/MSSkowron/GoMicroPriceFetcher/types"
	"github.com/sirupsen/logrus"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type RESTServer struct {
	listenAddr string
	svc        PriceFetcher
}

func NewRESTServer(a string, s PriceFetcher) *RESTServer {
	return &RESTServer{
		listenAddr: a,
		svc:        s,
	}
}

func (s *RESTServer) Run() {
	http.HandleFunc("/", makeHTTPHandler(s.handleFetchPrice))

	if err := http.ListenAndServe(s.listenAddr, nil); err != nil {
		log.Fatalln(err)
	}
}

func makeHTTPHandler(apiFn APIFunc) http.HandlerFunc {
	ctx := context.Background()

	ctx = context.WithValue(ctx, "requestID", rand.Intn(100000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			if err := writeJSON(w, http.StatusInternalServerError, map[string]any{"error": err.Error()}); err != nil {
				logrus.Errorln(err)
			}
		}
	}
}

func (s *RESTServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceResponse := &types.PriceResponse{
		Ticker: ticker,
		Price:  price,
	}

	return writeJSON(w, http.StatusOK, priceResponse)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
