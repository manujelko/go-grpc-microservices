package main

import (
	"net"

	"github.com/hashicorp/go-hclog"
	currencypb "github.com/manujelko/go-grpc-microservices/pkg/protobufs/currency"
	"github.com/manujelko/go-grpc-microservices/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrency(log)
	currencypb.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)
	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen")
	}
	gs.Serve(l)
}
