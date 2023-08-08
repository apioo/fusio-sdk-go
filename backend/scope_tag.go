// ScopeTag automatically generated by SDKgen please do not edit this file manually
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

type ScopeTag struct {
	internal *sdkgen.TagAbstract
}

// Delete
func (client *ScopeTag) Delete(scopeId string) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["scope_id"] = scopeId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/scope/$scope_id<[0-9]+|^~>", pathParams))
	if err != nil {
		return Message{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return Message{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Message{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 410:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// Update
func (client *ScopeTag) Update(scopeId string, payload ScopeUpdate) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["scope_id"] = scopeId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/scope/$scope_id<[0-9]+|^~>", pathParams))
	if err != nil {
		return Message{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return Message{}, errors.New("could not marshal provided JSON data")
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("PUT", u.String(), reqBody)
	if err != nil {
		return Message{}, errors.New("could not create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Message{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 410:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// Get
func (client *ScopeTag) Get(scopeId string) (Scope, error) {
	pathParams := make(map[string]interface{})
	pathParams["scope_id"] = scopeId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/scope/$scope_id<[0-9]+|^~>", pathParams))
	if err != nil {
		return Scope{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Scope{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Scope{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Scope{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Scope
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Scope{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Scope{}, errors.New("could not unmarshal JSON response")
		}

		return Scope{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Scope{}, errors.New("could not unmarshal JSON response")
		}

		return Scope{}, &MessageException{
			Payload: response,
		}
	case 410:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Scope{}, errors.New("could not unmarshal JSON response")
		}

		return Scope{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Scope{}, errors.New("could not unmarshal JSON response")
		}

		return Scope{}, &MessageException{
			Payload: response,
		}
	default:
		return Scope{}, errors.New("the server returned an unknown status code")
	}
}

// GetCategories
func (client *ScopeTag) GetCategories() (ScopeCategories, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/scope/categories", pathParams))
	if err != nil {
		return ScopeCategories{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return ScopeCategories{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return ScopeCategories{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ScopeCategories{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response ScopeCategories
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ScopeCategories{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ScopeCategories{}, errors.New("could not unmarshal JSON response")
		}

		return ScopeCategories{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ScopeCategories{}, errors.New("could not unmarshal JSON response")
		}

		return ScopeCategories{}, &MessageException{
			Payload: response,
		}
	default:
		return ScopeCategories{}, errors.New("the server returned an unknown status code")
	}
}

// Create
func (client *ScopeTag) Create(payload ScopeCreate) (Message, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/scope", pathParams))
	if err != nil {
		return Message{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return Message{}, errors.New("could not marshal provided JSON data")
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", u.String(), reqBody)
	if err != nil {
		return Message{}, errors.New("could not create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Message{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// GetAll
func (client *ScopeTag) GetAll(startIndex int, count int, search string) (ScopeCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	u, err := url.Parse(client.internal.Parser.Url("/backend/scope", pathParams))
	if err != nil {
		return ScopeCollection{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return ScopeCollection{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return ScopeCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ScopeCollection{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response ScopeCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ScopeCollection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ScopeCollection{}, errors.New("could not unmarshal JSON response")
		}

		return ScopeCollection{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ScopeCollection{}, errors.New("could not unmarshal JSON response")
		}

		return ScopeCollection{}, &MessageException{
			Payload: response,
		}
	default:
		return ScopeCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewScopeTag(httpClient *http.Client, parser *sdkgen.Parser) *ScopeTag {
	return &ScopeTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
