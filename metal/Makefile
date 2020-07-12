protos:
	protoc -I protos/ protos/metal.proto --go_out=plugins=grpc:protos/metal

clean:
	go mod download
	go mod tidy

build:
	docker build -t metal-service .

run:
	docker run -it -p 10502:10502 metal-service

rund:
	docker run -d -p 10502:10502 metal-service
