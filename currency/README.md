# Currency service
This service provides current exchange rates of two currencies.

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

### Run
The service can be run isolated:
```bash
[/currency] $ make build       # uses docker engine
[/currency] $ make run

docker run -it -p 10501:10501 --name currencysrv --rm currency-service
[CURRENCY SERVICE] 2020/07/14 07:31:21 Currency service is running (active)
```

Now you should be able to dial the service on port 10501 (by default). The service supports reflection. Try use grpcurl in another terminal:
```bash
$ grpcurl --plaintext -d '{"Base":"USD","Destination":"EUR"}' localhost:10501 currency.Currency.GetRate
{
    "Rate": 0.8826904
}

$ grpcurl --plaintext -d '{"Base":"CAD","Destination":"RUB"}' localhost:10501 currency.Currency.GetRate
{
    "Rate": 52.15866
}
```

Notice the log messages:
```bash
[CURRENCY SERVICE] 2020/07/14 07:31:21 Currency service is running (active)
[CURRENCY SERVICE] 2020/07/14 07:32:59 Handling GetRate; Base: USD, Destination: EUR
[CURRENCY SERVICE] 2020/07/14 07:34:56 Handling GetRate; Base: CAD, Destination: RUB
```

### Directory structure
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

### Test output
```bash
[/currency] $ go test -cover ./...
ok      github.com/chutified/metal-price/currency/config        0.006s  coverage: 100.0% of statements
ok      github.com/chutified/metal-price/currency/service       2.406s  coverage: 100.0% of statements
ok      github.com/chutified/metal-price/currency/service/data  (cached)        coverage: 86.8% of statements
ok      github.com/chutified/metal-price/currency/service/server        0.537s  coverage: 100.0% of statements
```
