// BackendCronjobByCronjobIdResource automatically generated by SDKgen please do not edit this file manually
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

type BackendCronjobByCronjobIdResource struct {
	url    string
	client *http.Client
}

// BackendActionCronjobGet
func (resource BackendCronjobByCronjobIdResource) BackendActionCronjobGet() (Cronjob, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return Cronjob{}, errors.New("could not parse url")
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return Cronjob{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return Cronjob{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Cronjob{}, errors.New("could not read response body")
	}

	var response Cronjob

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return Cronjob{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

// BackendActionCronjobUpdate
func (resource BackendCronjobByCronjobIdResource) BackendActionCronjobUpdate(data CronjobUpdate) (Message, error) {
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

// BackendActionCronjobDelete
func (resource BackendCronjobByCronjobIdResource) BackendActionCronjobDelete() (Message, error) {
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

func NewBackendCronjobByCronjobIdResource(cronjobId string, resource *sdkgen.Resource) *BackendCronjobByCronjobIdResource {
	return &BackendCronjobByCronjobIdResource{
		url:    resource.BaseUrl + "/backend/cronjob/" + cronjobId + "",
		client: resource.HttpClient,
	}
}
