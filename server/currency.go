package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/go-hclog"
	protos "github.com/sibis/currency-converter/protos/currency"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

type Conversion struct {
	Rates map[string]float32 `json:"rates"`
	Base  string             `json:"base"`
	Date  string             `json:"date"`
}

func (c *Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {

	c.log.Info("Handle Generate", "base", rr.GetBase(), "destination", rr.GetDestination())
	url := fmt.Sprintf("https://api.exchangeratesapi.io/latest?symbols=%v&base=%v", rr.GetBase(), rr.GetDestination())
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Err : ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var currency Conversion
	err = json.Unmarshal(body, &currency)
	str := fmt.Sprintf("%v", rr.GetBase())
	return &protos.RateResponse{Rate: currency.Rates[str]}, nil
}
