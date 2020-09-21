# Getting started

Mundipagg API

mundiapi requires a Go version with [Modules](https://github.com/golang/go/wiki/Modules) support and
uses import versioning. So please make sure to initialize a Go module before installing mundiapi:

```shell
go mod init github.com/my/repo
go get github.com/mundipagg/mundiapi
```

Import:

```go
import "github.com/mundipagg/mundiapi"
```

## Quickstart

```go
import (
    "github.com/mundipagg/mundiapi"
)

func ExampleNewClient() {
    client := mundiapi.NewClient("username", "password")

    resp, err := client.Charges().CreateCharge(&CreateChargeRequest{
        Amount:     100,
        Customer:   CreateCustomerRequest{
            Name:     "João da Silva", 
            Email:    "joao@gmail.com.br",
        },
    }, nil)

    fmt.Println(resp, err)
    // Output: GetChargeResponse <nil>
}
```
