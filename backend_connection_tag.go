// BackendConnectionTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package fusio

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/apioo/sdkgen-go"
	"io"
	"net/http"
	"net/url"
)

type BackendConnectionTag struct {
	internal *sdkgen.TagAbstract
}

// GetIntrospectionForEntity
func (client *BackendConnectionTag) GetIntrospectionForEntity(connectionId string, entity string) (BackendConnectionIntrospectionEntity, error) {
	pathParams := make(map[string]interface{})
	pathParams["connection_id"] = connectionId
	pathParams["entity"] = entity

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id<[0-9]+|^~>/introspection/:entity", pathParams))
	if err != nil {
		return BackendConnectionIntrospectionEntity{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendConnectionIntrospectionEntity{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendConnectionIntrospectionEntity{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendConnectionIntrospectionEntity{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendConnectionIntrospectionEntity
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionIntrospectionEntity{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionIntrospectionEntity{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnectionIntrospectionEntity{}, &CommonMessageException{
			Payload: response,
		}
	case 404:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionIntrospectionEntity{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnectionIntrospectionEntity{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionIntrospectionEntity{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnectionIntrospectionEntity{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendConnectionIntrospectionEntity{}, errors.New("the server returned an unknown status code")
	}
}

// GetIntrospection
func (client *BackendConnectionTag) GetIntrospection(connectionId string) (BackendConnectionIntrospectionEntities, error) {
	pathParams := make(map[string]interface{})
	pathParams["connection_id"] = connectionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id<[0-9]+|^~>/introspection", pathParams))
	if err != nil {
		return BackendConnectionIntrospectionEntities{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendConnectionIntrospectionEntities{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendConnectionIntrospectionEntities{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendConnectionIntrospectionEntities{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendConnectionIntrospectionEntities
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionIntrospectionEntities{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionIntrospectionEntities{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnectionIntrospectionEntities{}, &CommonMessageException{
			Payload: response,
		}
	case 404:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionIntrospectionEntities{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnectionIntrospectionEntities{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionIntrospectionEntities{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnectionIntrospectionEntities{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendConnectionIntrospectionEntities{}, errors.New("the server returned an unknown status code")
	}
}

// GetRedirect
func (client *BackendConnectionTag) GetRedirect(connectionId string) (CommonMessage, error) {
	pathParams := make(map[string]interface{})
	pathParams["connection_id"] = connectionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id<[0-9]+|^~>/redirect", pathParams))
	if err != nil {
		return CommonMessage{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return CommonMessage{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return CommonMessage{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CommonMessage{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return CommonMessage{}, errors.New("the server returned an unknown status code")
	}
}

// Delete
func (client *BackendConnectionTag) Delete(connectionId string) (CommonMessage, error) {
	pathParams := make(map[string]interface{})
	pathParams["connection_id"] = connectionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id<[0-9]+|^~>", pathParams))
	if err != nil {
		return CommonMessage{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return CommonMessage{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return CommonMessage{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CommonMessage{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 404:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 410:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return CommonMessage{}, errors.New("the server returned an unknown status code")
	}
}

// Update
func (client *BackendConnectionTag) Update(connectionId string, payload BackendConnectionUpdate) (CommonMessage, error) {
	pathParams := make(map[string]interface{})
	pathParams["connection_id"] = connectionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id<[0-9]+|^~>", pathParams))
	if err != nil {
		return CommonMessage{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return CommonMessage{}, errors.New("could not marshal provided JSON data")
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("PUT", u.String(), reqBody)
	if err != nil {
		return CommonMessage{}, errors.New("could not create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return CommonMessage{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CommonMessage{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 404:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 410:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return CommonMessage{}, errors.New("the server returned an unknown status code")
	}
}

// Get
func (client *BackendConnectionTag) Get(connectionId string) (BackendConnection, error) {
	pathParams := make(map[string]interface{})
	pathParams["connection_id"] = connectionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id<[0-9]+|^~>", pathParams))
	if err != nil {
		return BackendConnection{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendConnection{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendConnection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendConnection{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendConnection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 404:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnection{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnection{}, &CommonMessageException{
			Payload: response,
		}
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnection{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnection{}, &CommonMessageException{
			Payload: response,
		}
	case 410:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnection{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnection{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnection{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnection{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendConnection{}, errors.New("the server returned an unknown status code")
	}
}

// GetForm
func (client *BackendConnectionTag) GetForm(class string) (CommonFormContainer, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})
	queryParams["class"] = class

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/form", pathParams))
	if err != nil {
		return CommonFormContainer{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return CommonFormContainer{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return CommonFormContainer{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CommonFormContainer{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response CommonFormContainer
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonFormContainer{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonFormContainer{}, errors.New("could not unmarshal JSON response")
		}

		return CommonFormContainer{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonFormContainer{}, errors.New("could not unmarshal JSON response")
		}

		return CommonFormContainer{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return CommonFormContainer{}, errors.New("the server returned an unknown status code")
	}
}

// GetClasses
func (client *BackendConnectionTag) GetClasses() (BackendConnectionIndex, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/list", pathParams))
	if err != nil {
		return BackendConnectionIndex{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendConnectionIndex{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendConnectionIndex{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendConnectionIndex{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendConnectionIndex
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionIndex{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionIndex{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnectionIndex{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionIndex{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnectionIndex{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendConnectionIndex{}, errors.New("the server returned an unknown status code")
	}
}

// Create
func (client *BackendConnectionTag) Create(payload BackendConnectionCreate) (CommonMessage, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection", pathParams))
	if err != nil {
		return CommonMessage{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return CommonMessage{}, errors.New("could not marshal provided JSON data")
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", u.String(), reqBody)
	if err != nil {
		return CommonMessage{}, errors.New("could not create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return CommonMessage{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CommonMessage{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return CommonMessage{}, errors.New("could not unmarshal JSON response")
		}

		return CommonMessage{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return CommonMessage{}, errors.New("the server returned an unknown status code")
	}
}

// GetAll
func (client *BackendConnectionTag) GetAll(startIndex int, count int, search string) (BackendConnectionCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection", pathParams))
	if err != nil {
		return BackendConnectionCollection{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendConnectionCollection{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendConnectionCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendConnectionCollection{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendConnectionCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionCollection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionCollection{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnectionCollection{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendConnectionCollection{}, errors.New("could not unmarshal JSON response")
		}

		return BackendConnectionCollection{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendConnectionCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewBackendConnectionTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendConnectionTag {
	return &BackendConnectionTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
