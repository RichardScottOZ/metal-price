protos:
	protoc -I protos/ protos/metal.proto --go_out=plugins=grpc:protos/metal

clean:
	go mod download
	go mod tidy

build:
	docker build -t metal-service .

run:
	docker run -it --name metalsrv --rm metal-service

rund:
	docker run -d --name metalsrv --rm metal-service
