# Metal Price
Metal Price is a REST API which provides the current prices of the 6 precius metals in 33 different currencies.

The project uses gRPC in a microservices architecture. All services are containerized in Docker containers and the whole application can be easily run with the multi-container tool Docker Compose (<a href="https://github.com/chutified/metal-price/blob/master/docker-compose.yml">docker-compose.yml</a>).

## Services:
#### <a href="https://github.com/chutified/metal-price/tree/master/currency">Currency</a>
- <a href="https://github.com/chutified/metal-price/tree/master/currency#supported-currencies">Supported currencies</a>
- <a href="https://github.com/chutified/metal-price/tree/master/currency#directory-structure">Directory structure</a>
- <a href="https://github.com/chutified/metal-price/tree/master/currency#test-output">Test coverage</a>
- <a href="https://github.com/chutified/metal-price/blob/master/currency/Dockerfile">Dockerfile</a>

#### <a href="https://github.com/chutified/metal-price/tree/master/metal">Metal</a>
- <a href="https://github.com/chutified/metal-price/tree/master/metal#supported-precious-metals">Supported precious metals</a>
- <a href="https://github.com/chutified/metal-price/tree/master/metal#directory-structure">Drectory structure</a>
- <a href="https://github.com/chutified/metal-price/tree/master/metal#test-output">Test coverage</a>
- <a href="https://github.com/chutified/metal-price/blob/master/metal/Dockerfile">Dockerfile</a>

## Installation

#### Requirements
- Git
- Docker Engine (<a href="https://docs.docker.com/engine/install/" target="_blank">install</a>)
- Docker Compose (<a href="https://docs.docker.com/compose/install/" target="_blank">install</a>)

```bash
$ git clone https://github.com/chutified/metal-price.git

$ make build      # build or rebuild the service
$ make run        # start the docker containers

$ curl localhost:3001/ping
{"message":"pong"}
```

### API sources
- European Central Bank: <a href="https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml" target="_blank">ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml</a>
- Money Metals: <a href="https://www.moneymetals.com/api/spot-prices.json" target="_blank">moneymetals.com/api/spot-prices.json</a>

## Endpoints
| **Path** | **Response** |
|----------|------|
| `host:3001/ping`  | `OK` if the server is running |
| `host:3001/i/{metal}`  | current `{metal}` price per `oz` in `USD` |
| `host:3001/i/{metal}/{currency}`  | current `{metal}` price per `oz` in `{currency}` |
| `host:3001/i/{metal}/{currency}/{weight-unit}`  | current `{metal}` price per `{weight_unit}` in `{currency}` |

### Usage
```bash
$ make build      # build or rebuild the service
$ make run        # start the docker containers

docker-compose -f docker-compose.yml -p metal-pricer up

Starting metal-pricer_metalsrv_1    ... done
Starting metal-pricer_currencysrv_1 ... done
Starting metal-pricer_metal_price_1 ... done
Attaching to metal-pricer_currencysrv_1, metal-pricer_metalsrv_1, metal-pricer_metal_price_1

currencysrv_1  | [CURRENCY SERVICE] 2020/07/14 07:46:17 Currency service is running (active)
metalsrv_1     | [METAL SERVICE] 2020/07/14 07:46:17 Metal service is running (active)
metal_price_1  | [SERVER] 2020/07/14 07:46:18 Listening and serving HTTP on port 3001
```

### Examples
Run in another terminal.

#### host:3001/i/ *{metal}* :
```sh
$ curl localhost:3001/i/rhodium
{
    "metal": "rhodium",
    "price": 8100,
    "currency": "USD",
    "unit": "oz"
}
```

#### host:3001/i/ *{metal}* / *{currency}* :
```sh
$ curl localhost:3001/i/au/cad
{
    "metal": "gold",
    "price": 2455.86,
    "currency": "CAD",
    "unit": "oz"
}
```

#### host:3001/i/ *{metal}* / *{currency}* / *{weight-unit}* :
```sh
$ curl localhost:3001/i/ag/czk/kg
{
    "metal": "silver",
    "price": 15922.17,
    "currency": "CZK",
    "unit": "kg"
}
```  

Notice the log messages:
```bash
currencysrv_1  | [CURRENCY SERVICE] 2020/07/14 07:46:17 Currency service is running (active)
metalsrv_1     | [METAL SERVICE] 2020/07/14 07:46:17 Metal service is running (active)
metal_price_1  | [SERVER] 2020/07/14 07:46:18 Listening and serving HTTP on port 3001
metalsrv_1     | [METAL SERVICE] 2020/07/14 07:48:01 Handling GetPrice; Material: rhodium
metal_price_1  | [GIN] 2020/07/14 - 07:48:01 | 200 |  1.261795294s |      172.21.0.1 | GET      "/i/rhodium"
currencysrv_1  | [CURRENCY SERVICE] 2020/07/14 07:48:27 Handling GetRate; Base: USD, Destination: CAD
metalsrv_1     | [METAL SERVICE] 2020/07/14 07:48:28 Handling GetPrice; Material: gold
metal_price_1  | [GIN] 2020/07/14 - 07:48:28 | 200 |  986.333154ms |      172.21.0.1 | GET      "/i/au/cad"
currencysrv_1  | [CURRENCY SERVICE] 2020/07/14 07:48:30 Handling GetRate; Base: USD, Destination: CZK
metalsrv_1     | [METAL SERVICE] 2020/07/14 07:48:31 Handling GetPrice; Material: silver
metal_price_1  | [GIN] 2020/07/14 - 07:48:31 | 200 |   861.97045ms |      172.21.0.1 | GET      "/i/ag/czk/kg"
```

Use \<Ctrl-C\> to gracefully stop the server and all services:
```bash
Stopping metal-pricer_metal_price_1 ... done
Stopping metal-pricer_currencysrv_1 ... done
Stopping metal-pricer_metalsrv_1    ... done
```

### API documentation
Swagger 2.0: <a href="https://github.com/chutified/metal-price/blob/master/api-server/docs/swagger.json">swagger.json</a>

Run the service and visit <a href="http://localhost:3001/swagger/index.html" target="_blank">localhost:3001/swagger/index.html</a>.

## Directory structure
```
 /
 ├── api-server
 ├── currency
 ├── metal
 ├── docker-compose.yml
 ├── Makefile
 └── README.md
```
