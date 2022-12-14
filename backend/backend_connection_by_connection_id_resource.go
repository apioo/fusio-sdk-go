// BackendConnectionByConnectionIdResource automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/apioo/sdkgen-go"
	"io"
	"net/http"
	"net/url"
)

type BackendConnectionByConnectionIdResource struct {
	url    string
	client *http.Client
}

// BackendActionConnectionGet
func (resource BackendConnectionByConnectionIdResource) BackendActionConnectionGet() (Connection, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return Connection{}, errors.New("could not parse url")
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return Connection{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return Connection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Connection{}, errors.New("could not read response body")
	}

	var response Connection

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return Connection{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

// BackendActionConnectionUpdate
func (resource BackendConnectionByConnectionIdResource) BackendActionConnectionUpdate(data ConnectionUpdate) (Message, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return Message{}, errors.New("could not parse url")
	}

	raw, err := json.Marshal(data)
	if err != nil {
		return Message{}, errors.New("could not marshal provided JSON data")
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("PUT", url.String(), reqBody)
	if err != nil {
		return Message{}, errors.New("could not create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := resource.client.Do(req)
	if err != nil {
		return Message{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, errors.New("could not read response body")
	}

	var response Message

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return Message{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

// BackendActionConnectionDelete
func (resource BackendConnectionByConnectionIdResource) BackendActionConnectionDelete() (Message, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return Message{}, errors.New("could not parse url")
	}

	req, err := http.NewRequest("DELETE", url.String(), nil)
	if err != nil {
		return Message{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return Message{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, errors.New("could not read response body")
	}

	var response Message

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return Message{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

func NewBackendConnectionByConnectionIdResource(connectionId string, resource *sdkgen.Resource) *BackendConnectionByConnectionIdResource {
	return &BackendConnectionByConnectionIdResource{
		url:    resource.BaseUrl + "/backend/connection/" + connectionId + "",
		client: resource.HttpClient,
	}
}
