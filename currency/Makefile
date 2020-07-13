protos:
	protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency

clean:
	go mod download
	go mod tidy

# DOCKER RUN (isolated)

build:
	docker build -t currency-service .

run:
	docker run -it -p 10501:10501 --name currencysrv --rm currency-service

rund:
	docker run -d -p 10501:10501 --name currencysrv --rm currency-service
