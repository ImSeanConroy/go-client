# go-client
Go client library for custom Terraform Provider

## Getting Started

### Installing

To start using `go-client`, install Go and the following command:

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
	client, err := client.NewClient("http://localhost:3000", token)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	res, _ := client.Get("/tasks")
	fmt.Println(res.Get("0.description").String())
}
```

This will print something like:

```
Create go client
```

## Contributing

Contributions are welcome. Please open an issue or submit a pull request for any enhancements or bug fixes.

## License

This project is Distributed under the MIT License - see the [LICENSE](LICENSE) file for information.

## Support

If you are having problems, please let me know by [raising a new issue](https://github.com/ImSeanConroy/go-client/issues/new/choose).