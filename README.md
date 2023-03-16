# Service Binding Go Client for k8s Service Binding

<a href="https://github.com/RHEcosystemAppEng/sbo-go-library/actions?query=workflow%3Aunit-tests"><img alt="sbo-go-library unit tests status" src="https://github.com/RHEcosystemAppEng/sbo-go-library/workflows/unit-tests/badge.svg"></a>
[![Go Report Card](https://goreportcard.com/badge/github.com/RHEcosystemAppEng/sbo-go-library)](https://goreportcard.com/report/github.com/RHEcosystemAppEng/sbo-go-library) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/RHEcosystemAppEng/sbo-go-library.svg)](https://github.com/RHEcosystemAppEng/sbo-go-library/)

### Installation

```shell
go get -u github.com/RHEcosystemAppEng/sbo-go-library
```

### Usage

Have your connection properties saved under the `SERVICE_BINDING_ROOT=./bindings`

```shell
e.g.:
bindings
└── crdb
    ├── database
    ├── host
    ├── options
    ├── password
    ├── port
    ├── provider
    ├── sslmode
    ├── sslrootcert
    ├── type
    └── username
```

Add the client in your code

```go
import github.com/RHEcosystemAppEng/sbo-go-library/pkg/binding/convert
```

Get a MongoDB Connection string
```
connString, err := convert.GetMongoDBConnectionString()

# sample value for connString
mongodb+srv://someUserName:password123+isItAGoodPassword+7%40%25%7C%3F%5EB6@mongodb0.example.com:11010/random-db?retryWrites=true&w=majority
```

Get a PostgreSQL Connection string
```
connString, err := convert.GetPostgreSQLConnectionString()

# sample value for connString
postgresql://remote-user:%5C@aws.postgresql.com:5432/cloud-postgresql-DB?sslmode=disable&options=-c%20search_path=keyword%20-c%20geqo=off
```
  
run locally
```
SERVICE_BINDING_ROOT=bindings go run ./cmd/<main.go>
```

### Relevant info.

* [CockroachDB Cloud Sample Application](https://github.com/myeung18/cockroachdb-go-quickstart)
* [Service Binding Specification](https://github.com/k8s-service-bindings/spec)

