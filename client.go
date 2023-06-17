package fusio

import (
	"github.com/apioo/fusio-sdk-go/backend"
	"github.com/apioo/fusio-sdk-go/consumer"
	"github.com/apioo/sdkgen-go"
)

type Client struct {
	BaseUrl     string
	Credentials sdkgen.CredentialsInterface
}

func NewClient(baseUrl string, clientId string, clientSecret string, tokenStore sdkgen.TokenStoreInterface, scopes []string) *Client {
	var credentials = sdkgen.OAuth2{
		ClientId:         clientId,
		ClientSecret:     clientSecret,
		TokenUrl:         baseUrl + "/authorization/token",
		AuthorizationUrl: "",
		TokenStore:       tokenStore,
		Scopes:           scopes,
	}

	return &Client{
		BaseUrl:     baseUrl,
		Credentials: credentials,
	}
}

func (client *Client) Backend() (*backend.Client, error) {
	var backendClient, err = backend.NewClient(client.BaseUrl, client.Credentials)
	if err != nil {
		return &backend.Client{}, err
	}
	return backendClient, nil
}

func (client *Client) Consumer() (*consumer.Client, error) {
	var consumerClient, err = consumer.NewClient(client.BaseUrl, client.Credentials)
	if err != nil {
		return &consumer.Client{}, err
	}
	return consumerClient, nil
}
