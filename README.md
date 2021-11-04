# Service Binding Go Client for Service Binding Operator

<a href="https://github.com/RHEcosystemAppEng/sbo-go-library/actions?query=workflow%3Aunit-tests"><img alt="service-binding-client unit tests status" src="https://github.com/RHEcosystemAppEng/sbo-go-library/workflows/unit-tests/badge.svg"></a>
[![Go Report Card](https://goreportcard.com/badge/github.com/RHEcosystemAppEng/sbo-go-library)](https://goreportcard.com/report/github.com/RHEcosystemAppEng/sbo-go-library) [![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/RHEcosystemAppEng/sbo-go-library.svg)](https://github.com/RHEcosystemAppEng/sbo-go-library/)



Install

```shell
go get -u github.com/RHEcosystemAppEng/sbo-go-library
```

then include the client in your code
```go
import (
    github.com/RHEcosystemAppEng/sbo-go-library/pkg/binding/convert
)

// call
connString, err := convert.GetMongoDBConnectionString()
if err != nil {
    panic(err)
}
fmt.Println(connString)
connString, err = convert.GetPostgreSQLConnectionString()
if err != nil {
    panic(err)
}
fmt.Println(connString))
```
  
run locally
```
SERVICE_BINDING_ROOT=bindings go run ./cmd/<main.go>
```
