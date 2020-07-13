protos:
	protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency

clean:
	go mod download
	go mod tidy

build:
	docker build -t currency-service .

run:
	docker run -it --name currencysrv --rm currency-service

rund:
	docker run -d --name currencysrv --rm currency-service
