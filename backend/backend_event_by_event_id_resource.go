// BackendEventByEventIdResource automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

import (
	"bytes"

	"errors"
	"github.com/apioo/sdkgen-go"

	"net/http"
	"net/url"
)

type BackendEventByEventIdResource struct {
	url    string
	client *http.Client
}

// BackendActionEventGet
func (resource BackendEventByEventIdResource) BackendActionEventGet() (Event, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return Event{}, errors.New("could not parse url")
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return Event{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return Event{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Event{}, errors.New("could not read response body")
	}

	var response Event

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return Event{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

// BackendActionEventUpdate
func (resource BackendEventByEventIdResource) BackendActionEventUpdate(data EventUpdate) (Message, error) {
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

// BackendActionEventDelete
func (resource BackendEventByEventIdResource) BackendActionEventDelete() (Message, error) {
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

func NewBackendEventByEventIdResource(eventId string, resource *sdkgen.Resource) *BackendEventByEventIdResource {
	return &BackendEventByEventIdResource{
		url:    resource.BaseUrl + "/backend/event/" + eventId + "",
		client: resource.HttpClient,
	}
}
