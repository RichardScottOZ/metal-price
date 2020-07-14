# Metal service
This service provides current prices of precious metals.

### Supported precious metals
| **Symbol** | **Element** |
|------------|-------------|
| **Cu**  | copper | | **Ag**  | silver | | **Au**  | gold |
| **Pt**  | platium |
| **Pd**  | palladium |
| **Rh**  | rhodum |

*Both symbol and full element name can be used to select the metal.*


### Run
The service can  be run isolated:
```bash
[/metal] $ make build       # uses docker engine
[/metal] $ make run

docker run -it -p 10502:10502 --name metalsrv --rm metal-service
[METAL SERVICE] 2020/07/14 07:15:04 Metal service is running (active)
```

Now you should be able to dial the service on port 10502 (by default). The service supports reflection. Try use grpcurl in another terminal:
```bash
# grpcurl --plaintext -d '{"Metal":"gold"}' localhost:10502 metal.Metal.GetPrice
{
    "Price": 1809
}

$ grpcurl --plaintext -d '{"Metal":"platinum"}' localhost:10502 metal.Metal.GetPrice
{
    "Price": 845.4
}
```

Notice the log messages in the previous terminal:
```bash
[METAL SERVICE] 2020/07/14 07:15:04 Metal service is running (active)
[METAL SERVICE] 2020/07/14 07:20:24 Handling GetPrice; Material: gold
[METAL SERVICE] 2020/07/14 07:21:02 Handling GetPrice; Material: platinum
```

### Directory structure
```bash
 metal
 ├── config
 │   ├── config.go
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

### Test output
```bash
[/metal] $ go test -cover ./...
ok      github.com/chutified/metal-price/metal/config   0.002s  coverage: 100.0% of statements
ok      github.com/chutified/metal-price/metal/service  2.406s  coverage: 100.0% of statements
ok      github.com/chutified/metal-price/metal/service/data     (cached)        coverage: 89.5% of statements
ok      github.com/chutified/metal-price/metal/service/server   2.130s  coverage: 100.0% of statements
```
