package main

import (
	"github.com/apioo/fusio-sdk-go/sdk"
	"github.com/apioo/sdkgen-go"
	"reflect"
	"testing"
)

func TestClient(t *testing.T) {

	var credentials = sdkgen.HttpBearer{
		Token: "",
	}

	client, err := sdk.NewClient("https://api.acme.com", credentials)
	if err != nil {
		t.Error(err)
	}

	var backend = reflect.TypeOf(*client.Backend()).Name()
	if backend != "BackendTag" {
		t.Fatalf(`Client Backend not available, got: ` + backend)
	}

	var consumer = reflect.TypeOf(*client.Consumer()).Name()
	if consumer != "ConsumerTag" {
		t.Fatalf(`Client Consumer not available, got: ` + consumer)
	}

	var system = reflect.TypeOf(*client.System()).Name()
	if system != "SystemTag" {
		t.Fatalf(`Client System not available, got: ` + system)
	}

	var authorization = reflect.TypeOf(*client.Authorization()).Name()
	if authorization != "AuthorizationTag" {
		t.Fatalf(`Client Authorization not available, got: ` + authorization)
	}

}
