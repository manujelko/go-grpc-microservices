package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/manujelko/go-grpc-microservices/pkg/data"
	currencypb "github.com/manujelko/go-grpc-microservices/pkg/protobufs/currency"
	"github.com/manujelko/go-grpc-microservices/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()

	rates, err := data.NewExchangeRates(log)
	if err != nil {
		log.Error("Unable to generate rates")
		os.Exit(1)
	}

	cs := server.NewCurrency(rates, log)

	currencypb.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen")
	}

	gs.Serve(l)
}
