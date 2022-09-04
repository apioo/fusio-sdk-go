// ConsumerLogByLogIdResource automatically generated by SDKgen please do not edit this file manually
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

type ConsumerLogByLogIdResource struct {
	url    string
	client *http.Client
}

// ConsumerActionLogGet
func (resource ConsumerLogByLogIdResource) ConsumerActionLogGet(query CollectionQuery) (Log, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return Log{}, errors.New("could not parse url")
	}

	rawJson, err := json.Marshal(query)
	if err != nil {
		return Log{}, errors.New("could not marshall query")
	}

	parameters := url.Query()
	err = json.Unmarshal(rawJson, &parameters)
	url.RawQuery = parameters.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return Log{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return Log{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Log{}, errors.New("could not read response body")
	}

	var response Log

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return Log{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

func NewConsumerLogByLogIdResource(logId string, resource *sdkgen.Resource) *ConsumerLogByLogIdResource {
	return &ConsumerLogByLogIdResource{
		url:    resource.BaseUrl + "/consumer/log/" + logId + "",
		client: resource.HttpClient,
	}
}
