package main

import (
	"github.com/apioo/fusio-sdk-go/backend"
	"github.com/apioo/fusio-sdk-go/consumer"
	"github.com/apioo/sdkgen-go"
)

type Client struct {
	BaseUrl    string
	Scopes     []string
	TokenStore sdkgen.TokenStoreInterface
	Backend    *backend.Client
	Consumer   *consumer.Client
}

func NewClient(baseUrl string, clientId string, clientSecret string, tokenStore sdkgen.TokenStoreInterface, scopes []string) *Client {
	var credentials = sdkgen.ClientCredentials{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}

	return &Client{
		BaseUrl:    baseUrl,
		Scopes:     scopes,
		TokenStore: tokenStore,
		Backend:    backend.NewClient(baseUrl, credentials, tokenStore, scopes),
		Consumer:   consumer.NewClient(baseUrl, credentials, tokenStore, scopes),
	}
}

func main() {

	var store = &sdkgen.MemoryTokenStore{}
	var scopes = []string{"foo", "bar"}

	var foo = NewClient("", "", "", store, scopes)

	var data = backend.ActionCreate{
		Status: 1,
		Async:  false,
		Name:   "foo",
		Class:  "",
	}

	foo.Backend.BackendAction().GetBackendAction().BackendActionActionCreate(data)

}
