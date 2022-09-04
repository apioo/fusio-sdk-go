// ConsumerLogResource automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package consumer

import (
	"encoding/json"
	"errors"
	"github.com/apioo/sdkgen-go"
	"io"
	"net/http"
	"net/url"
)

type ConsumerLogResource struct {
	url    string
	client *http.Client
}

// ConsumerActionLogGetAll
func (resource ConsumerLogResource) ConsumerActionLogGetAll(query CollectionQuery) (LogCollection, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return LogCollection{}, errors.New("could not parse url")
	}

	rawJson, err := json.Marshal(query)
	if err != nil {
		return LogCollection{}, errors.New("could not marshall query")
	}

	parameters := url.Query()
	err = json.Unmarshal(rawJson, &parameters)
	url.RawQuery = parameters.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return LogCollection{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return LogCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return LogCollection{}, errors.New("could not read response body")
	}

	var response LogCollection

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return LogCollection{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

func NewConsumerLogResource(resource *sdkgen.Resource) *ConsumerLogResource {
	return &ConsumerLogResource{
		url:    resource.BaseUrl + "/consumer/log",
		client: resource.HttpClient,
	}
}
