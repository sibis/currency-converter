package server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	protos "github.com/sibis/currency-converter/protos/currency"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle Getrate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
