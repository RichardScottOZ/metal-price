.PHONY: clean build run rund protos

# UPDATE PROTO FILE

protos:
	protoc -I protos/ protos/metal.proto --go_out=plugins=grpc:protos/metal

# CLEAN UNWANTED DEPENDENCIES

clean:
	go mod download
	go mod tidy

# DOCKER RUN (isolated)

build:
	docker build -t metal-service .

run:
	docker run -it -p 10502:10502 --name metalsrv --rm metal-service

rund:
	docker run -d -p 10502:10502 --name metalsrv --rm metal-service
