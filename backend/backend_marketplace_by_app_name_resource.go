// BackendMarketplaceByAppNameResource automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

import (
	"bytes"

	"errors"
	"github.com/apioo/sdkgen-go"

	"net/http"
	"net/url"
)

type BackendMarketplaceByAppNameResource struct {
	url    string
	client *http.Client
}

// BackendActionMarketplaceGet
func (resource BackendMarketplaceByAppNameResource) BackendActionMarketplaceGet() (MarketplaceLocalApp, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return MarketplaceLocalApp{}, errors.New("could not parse url")
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return MarketplaceLocalApp{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return MarketplaceLocalApp{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return MarketplaceLocalApp{}, errors.New("could not read response body")
	}

	var response MarketplaceLocalApp

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return MarketplaceLocalApp{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

// BackendActionMarketplaceUpdate
func (resource BackendMarketplaceByAppNameResource) BackendActionMarketplaceUpdate() (Message, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return Message{}, errors.New("could not parse url")
	}

	req, err := http.NewRequest("PUT", url.String(), nil)
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

// BackendActionMarketplaceRemove
func (resource BackendMarketplaceByAppNameResource) BackendActionMarketplaceRemove() (Message, error) {
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

func NewBackendMarketplaceByAppNameResource(appName string, resource *sdkgen.Resource) *BackendMarketplaceByAppNameResource {
	return &BackendMarketplaceByAppNameResource{
		url:    resource.BaseUrl + "/backend/marketplace/" + appName + "",
		client: resource.HttpClient,
	}
}
