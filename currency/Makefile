proto-gen:
	protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency

run:
	go run main.go

clean:
	go mod download
	go mod tidy
