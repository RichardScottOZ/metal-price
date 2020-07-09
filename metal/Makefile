proto-gen:
	protoc -I protos/ protos/metal.proto --go_out=plugins=grpc:protos/metal

run:
	go run main.go

clean:
	go mod download
	go mod tidy
