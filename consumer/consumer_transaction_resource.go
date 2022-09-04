// ConsumerTransactionResource automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package consumer

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/apioo/sdkgen-go"
	"io"
	"net/http"
	"net/url"
)

type ConsumerTransactionResource struct {
	url    string
	client *http.Client
}

// ConsumerActionTransactionGetAll
func (resource ConsumerTransactionResource) ConsumerActionTransactionGetAll(query CollectionQuery) (TransactionCollection, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return TransactionCollection{}, errors.New("could not parse url")
	}

	rawJson, err := json.Marshal(query)
	if err != nil {
		return TransactionCollection{}, errors.New("could not marshall query")
	}

	parameters := url.Query()
	err = json.Unmarshal(rawJson, &parameters)
	url.RawQuery = parameters.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return TransactionCollection{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return TransactionCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return TransactionCollection{}, errors.New("could not read response body")
	}

	var response TransactionCollection

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return TransactionCollection{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

func NewConsumerTransactionResource(resource *sdkgen.Resource) *ConsumerTransactionResource {
	return &ConsumerTransactionResource{
		url:    resource.BaseUrl + "/consumer/transaction",
		client: resource.HttpClient,
	}
}
