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
