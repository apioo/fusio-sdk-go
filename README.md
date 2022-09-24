
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
	"github.com/apioo/fusio-sdk-go/backend"
	"github.com/apioo/sdkgen-go"
	"log"
)

func main() {
	var store = &sdkgen.MemoryTokenStore{}
	var scopes = []string{"backend"}

	var client = fusio.NewClient("https://demo.fusio-project.org", "test", "FRsNh1zKCXlB", store, scopes)

	collection, error := client.Backend.GetBackendRoutes().BackendActionRouteGetAll(backend.CollectionCategoryQuery{})

	if error != nil {
		log.Panic(error)
	}

	fmt.Println("Routes:")
	for _, v := range collection.Entry {
		fmt.Println("* " + v.Path)
	}
}

```
