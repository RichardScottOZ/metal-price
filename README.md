TODO:
    content, api, test, tree (3x), docker, api documentation, make commands, grpc (microservices) - grpcurl examples + results, instalation - make install, author, API sources, docker compose, makefile - install: git clone, make build, make run - prerequirements



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
