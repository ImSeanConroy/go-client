# go-client
Go client library for custom Terraform Provider

## Getting Started

### Installing

To start using `go-client`, install Go and `go get`:

`$ go get -u github.com/imseanconroy/go-client`

### Basic Usage

```go
package main

import (
	"fmt"

	"github.com/imseanconroy/go-client"
)

func main() {
	token := "YOUR_TOKEN"
	client := client.NewClient("http://localhost:3000", token)

	res, _ := client.Get("/tasks")
	fmt.Println(res.Get("0.description").String())
}
```

This will print something like:

```
Create go client
```