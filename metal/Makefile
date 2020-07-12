protos:
	protoc -I protos/ protos/metal.proto --go_out=plugins=grpc:protos/metal

run:
	go run main.go

clean:
	go mod download
	go mod tidy

docker-build:
	docker build -t metal-service .

docker-run:
	docker run -it -p 10502:10502 metal-service
