// ConsumerTransactionByTransactionIdResource automatically generated by SDKgen please do not edit this file manually
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

type ConsumerTransactionByTransactionIdResource struct {
	url    string
	client *http.Client
}

// ConsumerActionTransactionGet
func (resource ConsumerTransactionByTransactionIdResource) ConsumerActionTransactionGet() (Transaction, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return Transaction{}, errors.New("could not parse url")
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return Transaction{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return Transaction{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Transaction{}, errors.New("could not read response body")
	}

	var response Transaction

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return Transaction{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

func NewConsumerTransactionByTransactionIdResource(transactionId string, resource *sdkgen.Resource) *ConsumerTransactionByTransactionIdResource {
	return &ConsumerTransactionByTransactionIdResource{
		url:    resource.BaseUrl + "/consumer/transaction/" + transactionId + "",
		client: resource.HttpClient,
	}
}
