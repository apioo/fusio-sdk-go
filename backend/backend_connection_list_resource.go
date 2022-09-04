// BackendConnectionListResource automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

import (
	"encoding/json"
	"errors"
	"github.com/apioo/sdkgen-go"
	"io"
	"net/http"
	"net/url"
)

type BackendConnectionListResource struct {
	url    string
	client *http.Client
}

// BackendActionConnectionGetIndex
func (resource BackendConnectionListResource) BackendActionConnectionGetIndex() (ConnectionIndex, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return ConnectionIndex{}, errors.New("could not parse url")
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return ConnectionIndex{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return ConnectionIndex{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ConnectionIndex{}, errors.New("could not read response body")
	}

	var response ConnectionIndex

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return ConnectionIndex{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

func NewBackendConnectionListResource(resource *sdkgen.Resource) *BackendConnectionListResource {
	return &BackendConnectionListResource{
		url:    resource.BaseUrl + "/backend/connection/list",
		client: resource.HttpClient,
	}
}
