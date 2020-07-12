run:
	go run main.go

clean:
	go mod download
	go mod tidy

build:
	docker build -t metal-price-api .

run:
	docker run -it -p 8080:8080 metal-price-api

rund:
	docker run -d -p 8080:8080 metal-price-api
