// BackendActionListResource automatically generated by SDKgen please do not edit this file manually
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

type BackendActionListResource struct {
	url    string
	client *http.Client
}

// BackendActionActionGetIndex
func (resource BackendActionListResource) BackendActionActionGetIndex() (ActionIndex, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return ActionIndex{}, errors.New("could not parse url")
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return ActionIndex{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return ActionIndex{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ActionIndex{}, errors.New("could not read response body")
	}

	var response ActionIndex

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return ActionIndex{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

func NewBackendActionListResource(resource *sdkgen.Resource) *BackendActionListResource {
	return &BackendActionListResource{
		url:    resource.BaseUrl + "/backend/action/list",
		client: resource.HttpClient,
	}
}
