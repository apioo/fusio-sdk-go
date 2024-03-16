
// Client automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    
    
    
    "github.com/apioo/sdkgen-go"
    
    
    
    
)

type Client struct {
    internal *sdkgen.ClientAbstract
}

func (client *Client) Authorization() *AuthorizationTag {
    return NewAuthorizationTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *Client) System() *SystemTag {
    return NewSystemTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *Client) Consumer() *ConsumerTag {
    return NewConsumerTag(client.internal.HttpClient, client.internal.Parser)
}

func (client *Client) Backend() *BackendTag {
    return NewBackendTag(client.internal.HttpClient, client.internal.Parser)
}




func NewClient(baseUrl string, credentials sdkgen.CredentialsInterface) (*Client, error) {
	var client, err = sdkgen.NewClient(baseUrl, credentials)
    if err != nil {
        return &Client{}, err
    }

	return &Client{
		internal: client,
	}, nil
}

func Build(clientId string, clientSecret string, tokenStore sdkgen.TokenStoreInterface, scopes []string) (*Client, error) {
    var credentials = sdkgen.OAuth2{
        ClientId: clientId,
        ClientSecret: clientSecret,
        TokenUrl: "https://api.typehub.cloud/authorization/token",
        AuthorizationUrl: "",
        TokenStore: tokenStore,
        Scopes: scopes,
    }

    return NewClient("https://api.typehub.cloud/", credentials)
}

