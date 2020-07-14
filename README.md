TODO:
    content, api, test, docker, grpc (microservices), docker compose

## Installation

### Requirements
- Git
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

### Examples

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
$ curl localhost:3001/i/au/cad | jq

{
"metal": "gold",
"price": 2455.86,
"currency": "CAD",
"unit": "oz"
}
```


#### host:3001/i/ *{metal}* / *{currency}* / *{weight-unit}* :
```sh
$ curl localhost:3001/i/ag/czk/kg | jq

{
"metal": "silver",
"price": 15922.17,
"currency": "CZK",
"unit": "kg"
}
```  

### API documentation
Swagger 2.0: <a href="">swagger.json</a>
Run the service and visit: <a href="http://localhost:3001/swagger/index.html" target="_blank">localhost:3001/swagger/index.html</a>

### Supported weight units
| **Sign** | **Unit** |
|----------|----------|
| **oz**  | ounce |
| **lb**  | pound |
| **g**   | gram |
| **dkg** | decagram |
| **kg**  | kilogram |
| **t**   | ton |

*Both sign and unit name can be used to select the weight unit.*


### Supported precious metal
| **Symbol** | **Element** |
|------------|-------------|
| **Cu**  | copper |
| **Ag**  | silver |
| **Au**  | gold |
| **Pt**  | platium |
| **Pd**  | palladium |
| **Rh**  | rhodum |

*Both symbol and full element name can be used to select the metal.*


### Supported currencies
<table>
    <tr> <td>EUR</td> <td>CAD</td> <td>HKD</td> <td>ISK</td> <td>PHP</td> </tr>
    <tr> <td>DKK</td> <td>HUG</td> <td>CZK</td> <td>AUD</td> <td>RON</td> </tr>
    <tr> <td>SEK</td> <td>IDR</td> <td>INR</td> <td>BRL</td> <td>RUB</td> </tr>
    <tr> <td>HRK</td> <td>JPY</td> <td>THB</td> <td>CHF</td> <td>SGD</td> </tr>
    <tr> <td>PLN</td> <td>BGN</td> <td>TRY</td> <td>CNY</td> <td>NOK</td> </tr>
    <tr> <td>NZD</td> <td>ZAR</td> <td>USD</td> <td>MXN</td> <td>ILS</td> </tr>
    <tr> <td>GBP</td> <td>KRW</td> <td>MYR</td> </tr>
</table>


## Directory structure

### Root dir
```
 /
 ├── api-server
 ├── currency
 ├── metal
 ├── docker-compose.yml
 ├── Makefile
 └── README.md
```

### Web server
```bash
 api-server
 ├── app
 │   ├── handlers
 │   │   ├── handler.go
 │   │   ├── handlers_test.go
 │   │   ├── ping.go
 │   │   ├── price-mc.go
 │   │   ├── price-mcu.go
 │   │   ├── price-m.go
 │   │   ├── response-model.go
 │   │   └── routes.go
 │   ├── services
 │   │   ├── currency.go
 │   │   ├── currency_test.go
 │   │   ├── metal.go
 │   │   ├── metal_test.go
 │   │   ├── periodic-symbols.go
 │   │   ├── weightconv.go
 │   │   └── weightconv_test.go
 │   ├── app.go
 │   └── app_test.go
 ├── config
 │   ├── config.go
 │   └── config_test.go
 ├── docs
 │   ├── docs.go
 │   ├── swagger.json
 │   └── swagger.yaml
 ├── Dockerfile
 ├── go.mod
 ├── go.sum
 ├── main.go
 └── Makefile
```

### Currency service
```bash
 currency
 ├── config
 │   ├── config.go
 │   └── config_test.go
 ├── service
 │   ├── data
 │   │   ├── rates.go
 │   │   └── rates_test.go
 │   ├── protos
 │   │   ├── currency
 │   │   │   └── currency.pb.go
 │   │   └── currency.proto
 │   ├── server
 │   │   ├── currency.go
 │   │   └── currency_test.go
 │   ├── service.go
 │   └── service_test.go
 ├── Dockerfile
 ├── go.mod
 ├── go.sum
 ├── main.go
 └── Makefile
```

### Metal service
```bash
 metal ├── config │   ├── config.go
 │   └── config_test.go
 ├── service
 │   ├── data
 │   │   ├── prices.go
 │   │   └── prices_test.go
 │   ├── protos
 │   │   ├── metal
 │   │   │   └── metal.pb.go
 │   │   └── metal.proto
 │   ├── server
 │   │   ├── metal.go
 │   │   └── metal_test.go
 │   ├── service.go
 │   └── service_test.go
 ├── Dockerfile
 ├── go.mod
 ├── go.sum
 ├── main.go
 └── Makefile
```

## Testing
```bash
[/currency] $ go test -cover ./...
ok      github.com/chutified/metal-price/currency/config        0.006s  coverage: 100.0% of statements
ok      github.com/chutified/metal-price/currency/service       2.406s  coverage: 100.0% of statements
ok      github.com/chutified/metal-price/currency/service/data  (cached)        coverage: 86.8% of statements
ok      github.com/chutified/metal-price/currency/service/server        0.537s  coverage: 100.0% of statements
```

```bash
[/metal] $ go test -cover ./...
ok      github.com/chutified/metal-price/metal/config   0.002s  coverage: 100.0% of statements
ok      github.com/chutified/metal-price/metal/service  2.406s  coverage: 100.0% of statements
ok      github.com/chutified/metal-price/metal/service/data     (cached)        coverage: 89.5% of statements
ok      github.com/chutified/metal-price/metal/service/server   2.130s  coverage: 100.0% of statements
```

```bash
[/api-server] $ go test -cover ./...
ok      github.com/chutified/metal-price/api-server/app (cached)        coverage: 91.3% of statements
ok      github.com/chutified/metal-price/api-server/app/handlers        7.167s  coverage: 100.0% of statements
ok      github.com/chutified/metal-price/api-server/app/services        1.607s  coverage: 100.0% of statements
ok      github.com/chutified/metal-price/api-server/config      (cached)        coverage: 100.0% of statements
```
