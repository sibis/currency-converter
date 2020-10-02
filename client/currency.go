package main

import (
	"context"
	"fmt"

	protos "github.com/sibis/currency-converter/protos/currency"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8085", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cc := protos.NewCurrencyClient(conn)

	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["INR"]),
		Destination: protos.Currencies(protos.Currencies_value["USD"]),
	}
	resp, err := cc.GetRate(context.Background(), rr)

	if err != nil {
		fmt.Println("[error] getting new rate ", err)
	}
	fmt.Println("resp data : ", resp)
}
