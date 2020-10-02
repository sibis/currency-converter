package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	protos "github.com/sibis/currency-converter/protos/currency"
	"github.com/sibis/currency-converter/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// https://api.exchangeratesapi.io/latest?symbols=INR&base=USD
func main() {
	log := hclog.Default()
	cs := server.NewCurrency(log)

	gs := grpc.NewServer()

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Error("Unable to listen", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
