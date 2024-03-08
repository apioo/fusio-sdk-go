// BackendRoleTag automatically generated by SDKgen please do not edit this file manually
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

type BackendRoleTag struct {
	internal *sdkgen.TagAbstract
}

// Delete
func (client *BackendRoleTag) Delete(roleId string) (CommonMessage, error) {
	pathParams := make(map[string]interface{})
	pathParams["role_id"] = roleId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/role/$role_id<[0-9]+|^~>", pathParams))
	if err != nil {
		return CommonMessage{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return CommonMessage{}, err
	}

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

// Update
func (client *BackendRoleTag) Update(roleId string, payload BackendRoleUpdate) (CommonMessage, error) {
	pathParams := make(map[string]interface{})
	pathParams["role_id"] = roleId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/role/$role_id<[0-9]+|^~>", pathParams))
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
func (client *BackendRoleTag) Get(roleId string) (BackendRole, error) {
	pathParams := make(map[string]interface{})
	pathParams["role_id"] = roleId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/role/$role_id<[0-9]+|^~>", pathParams))
	if err != nil {
		return BackendRole{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendRole{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendRole{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendRole{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendRole
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendRole{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendRole{}, err
		}

		return BackendRole{}, &CommonMessageException{
			Payload: response,
		}
	case 404:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendRole{}, err
		}

		return BackendRole{}, &CommonMessageException{
			Payload: response,
		}
	case 410:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendRole{}, err
		}

		return BackendRole{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendRole{}, err
		}

		return BackendRole{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendRole{}, errors.New("the server returned an unknown status code")
	}
}

// Create
func (client *BackendRoleTag) Create(payload BackendRoleCreate) (CommonMessage, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/role", pathParams))
	if err != nil {
		return CommonMessage{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return CommonMessage{}, err
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", u.String(), reqBody)
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

// GetAll
func (client *BackendRoleTag) GetAll(startIndex int, count int, search string) (BackendRoleCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	u, err := url.Parse(client.internal.Parser.Url("/backend/role", pathParams))
	if err != nil {
		return BackendRoleCollection{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendRoleCollection{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendRoleCollection{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendRoleCollection{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendRoleCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendRoleCollection{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendRoleCollection{}, err
		}

		return BackendRoleCollection{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendRoleCollection{}, err
		}

		return BackendRoleCollection{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendRoleCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewBackendRoleTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendRoleTag {
	return &BackendRoleTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}