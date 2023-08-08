
# Fusio Go SDK

This is the official Fusio Go SDK, it helps to talk to the Fusio REST API.
Fusio is an open source API management system, more information at:
https://www.fusio-project.org

## Usage

The following example shows how you can get all registered routes at the backend.
A working example is also available at: https://github.com/apioo/fusio-sample-go-cli

```go
package main

import (
	"fmt"
	"github.com/apioo/fusio-sdk-go"
	"github.com/apioo/sdkgen-go"
	"log"
)

func main() {
	var store = &sdkgen.MemoryTokenStore{}
	var scopes = []string{"backend"}

	var client = fusio.NewClient("https://demo.fusio-project.org", "test", "FRsNh1zKCXlB", store, scopes)

	backend, err := client.Backend()
	if err != nil {
		log.Panic(err)
	}

	collection, err := backend.Operation().GetAll(0, 16, "")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Operations:")
	for _, v := range collection.Entry {
		fmt.Println("* " + v.HttpMethod + " " + v.HttpPath)
	}
}

```
