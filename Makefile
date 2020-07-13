.PHONY: build run

build:
	docker-compose -f docker-compose.yml -p metal-pricer build

run:
	docker-compose -f docker-compose.yml -p metal-pricer up
