// BackendConfigTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/apioo/sdkgen-go"
	"io"
	"net/http"
	"net/url"
)

type BackendConfigTag struct {
	internal *sdkgen.TagAbstract
}

// Update
func (client *BackendConfigTag) Update(configId string, payload BackendConfigUpdate) (CommonMessage, error) {
	pathParams := make(map[string]interface{})
	pathParams["config_id"] = configId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/config/$config_id<[0-9]+|^~>", pathParams))
	if err != nil {
		return CommonMessage{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return CommonMessage{}, err
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("PUT", u.String(), reqBody)
	if err != nil {
		return CommonMessage{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return CommonMessage{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CommonMessage{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, err
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, err
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 404:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, err
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 410:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, err
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, err
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return CommonMessage{}, errors.New("the server returned an unknown status code")
	}
}

// Get
func (client *BackendConfigTag) Get(configId string) (BackendConfig, error) {
	pathParams := make(map[string]interface{})
	pathParams["config_id"] = configId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/config/$config_id<[0-9]+|^~>", pathParams))
	if err != nil {
		return BackendConfig{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendConfig{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendConfig{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendConfig{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendConfig
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConfig{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConfig{}, err
		}

		return BackendConfig{}, &CommonMessageException{
			Payload: response,
		}
	case 404:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConfig{}, err
		}

		return BackendConfig{}, &CommonMessageException{
			Payload: response,
		}
	case 410:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConfig{}, err
		}

		return BackendConfig{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConfig{}, err
		}

		return BackendConfig{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendConfig{}, errors.New("the server returned an unknown status code")
	}
}

// GetAll
func (client *BackendConfigTag) GetAll(startIndex int, count int, search string) (BackendConfigCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	u, err := url.Parse(client.internal.Parser.Url("/backend/config", pathParams))
	if err != nil {
		return BackendConfigCollection{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendConfigCollection{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendConfigCollection{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendConfigCollection{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendConfigCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConfigCollection{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConfigCollection{}, err
		}

		return BackendConfigCollection{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConfigCollection{}, err
		}

		return BackendConfigCollection{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendConfigCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewBackendConfigTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendConfigTag {
	return &BackendConfigTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}