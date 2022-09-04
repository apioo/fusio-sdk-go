
# Fusio SDK

This Go library helps to talk to the Fusio (https://www.fusio-project.org/)
API. It uses the automatically generated definition of the Fusio backend.

## Usage

The following example shows how you can get a client for the backend to list
all routes at the backend. You can also checkout our sample go CLI [repository](https://github.com/apioo/fusio-sample-gocli).

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

	collection, error := client.Backend.BackendRoute().GetBackendRoutes().BackendActionRouteGetAll(backend.CollectionCategoryQuery{})

	if error != nil {
		log.Panic(error)
	}

	for _, v := range collection.Entry {
		fmt.Println("Route: " + v.Path)
	}
}

```
