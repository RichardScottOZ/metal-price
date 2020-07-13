clean:
	go mod download
	go mod tidy

build:
	docker build -t metal-price-api .

run:
	docker run -it -p 8080:8080 --name metal_price --rm metal-price-api

rund:
	docker run -d -p 8080:8080 --name metal_price --rm metal-price-api
