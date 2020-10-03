# currency-converter
GRPC based app to convert the current currency rates. This uses protobuf instead of json and works with HTTP 2.0 protocol.

Standalone microservice built for learning purpose accepts the base and target currency and returns the current conversion rate for it.


Setup:

proto file located under, `protos/currency.proto`

To generate the go output, 
run `make protos` (declaration can be found in Makefile)


Server declaration can be found under, `server/currency.go`

To run the server, `go run server/currency.go`


Test whether server responds using grpc curl,
`grpcurl --plaintext -d '{"Base":"INR", "Destination":"USD"}' localhost:8085 Currency.GetRate`

To make use of this microservice from other service or function, please look into client side implementation under, `client/currency.go` and to test the client with sample test, run `go run client/currency.go` (make sure the server is running to fetch the response)


