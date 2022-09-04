// BackendScopeResource automatically generated by SDKgen please do not edit this file manually
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

type BackendScopeResource struct {
	url    string
	client *http.Client
}

// BackendActionScopeGetAll
func (resource BackendScopeResource) BackendActionScopeGetAll(query CollectionCategoryQuery) (ScopeCollection, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return ScopeCollection{}, errors.New("could not parse url")
	}

	rawJson, err := json.Marshal(query)
	if err != nil {
		return ScopeCollection{}, errors.New("could not marshall query")
	}

	parameters := url.Query()
	err = json.Unmarshal(rawJson, &parameters)
	url.RawQuery = parameters.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return ScopeCollection{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return ScopeCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ScopeCollection{}, errors.New("could not read response body")
	}

	var response ScopeCollection

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return ScopeCollection{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

// BackendActionScopeCreate
func (resource BackendScopeResource) BackendActionScopeCreate(data ScopeCreate) (Message, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return Message{}, errors.New("could not parse url")
	}

	raw, err := json.Marshal(data)
	if err != nil {
		return Message{}, errors.New("could not marshal provided JSON data")
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", url.String(), reqBody)
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

func NewBackendScopeResource(resource *sdkgen.Resource) *BackendScopeResource {
	return &BackendScopeResource{
		url:    resource.BaseUrl + "/backend/scope",
		client: resource.HttpClient,
	}
}
