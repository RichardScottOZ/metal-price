# Web API

### Build-in service: Weight converter
Supported weight units:

| **Sign** | **Unit** |
|----------|----------|
| **oz**  | ounce |
| **lb**  | pound |
| **g**   | gram |
| **dkg** | decagram |
| **kg**  | kilogram |
| **t**   | ton |

*Both sign and unit name can be used to select the weight unit.*

### Directory structure
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

### Test output
```bash
[/api-server] $ go test -cover ./...
ok      github.com/chutified/metal-price/api-server/app (cached)        coverage: 91.3% of statements
ok      github.com/chutified/metal-price/api-server/app/handlers        7.167s  coverage: 100.0% of statements
ok      github.com/chutified/metal-price/api-server/app/services        1.607s  coverage: 100.0% of statements
ok      github.com/chutified/metal-price/api-server/config      (cached)        coverage: 100.0% of statements
```
