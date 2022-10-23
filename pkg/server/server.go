package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	currencypb "github.com/manujelko/go-grpc-microservices/pkg/protobufs/currency"
)

type Currency struct {
	log hclog.Logger
	currencypb.UnimplementedCurrencyServer
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l, currencypb.UnimplementedCurrencyServer{}}
}

func (c *Currency) GetRate(ctx context.Context, req *currencypb.GetRateRequest) (*currencypb.GetRateResponse, error) {
	c.log.Info("Handle GetRate", "base", req.GetBase(), "destination", req.GetDestination())

	return &currencypb.GetRateResponse{Rate: 0.5}, nil
}
