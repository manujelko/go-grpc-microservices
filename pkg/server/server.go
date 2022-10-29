package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/manujelko/go-grpc-microservices/pkg/data"
	currencypb "github.com/manujelko/go-grpc-microservices/pkg/protobufs/currency"
)

type Currency struct {
	rates *data.ExchangeRates
	log   hclog.Logger
	currencypb.UnimplementedCurrencyServer
}

func NewCurrency(r *data.ExchangeRates, l hclog.Logger) *Currency {
	return &Currency{r, l, currencypb.UnimplementedCurrencyServer{}}
}

func (c *Currency) GetRate(ctx context.Context, req *currencypb.GetRateRequest) (*currencypb.GetRateResponse, error) {
	c.log.Info("Handle GetRate", "base", req.GetBase(), "destination", req.GetDestination())

	rate, err := c.rates.GetRate(req.GetBase().String(), req.GetDestination().String())
	if err != nil {
		return nil, err
	}

	return &currencypb.GetRateResponse{Rate: rate}, nil
}
