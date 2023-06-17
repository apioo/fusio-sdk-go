// ActionTag automatically generated by SDKgen please do not edit this file manually
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

type ActionTag struct {
	internal *sdkgen.TagAbstract
}

// Delete
func (client *ActionTag) Delete(actionId string) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["action_id"] = actionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/action/$action_id&lt;[0-9]+|^~&gt;", pathParams))
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

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
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

	switch resp.StatusCode {
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// Update
func (client *ActionTag) Update(actionId string, payload ActionUpdate) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["action_id"] = actionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/action/$action_id&lt;[0-9]+|^~&gt;", pathParams))
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

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
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

	switch resp.StatusCode {
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// Get
func (client *ActionTag) Get(actionId string) (Action, error) {
	pathParams := make(map[string]interface{})
	pathParams["action_id"] = actionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/action/$action_id&lt;[0-9]+|^~&gt;", pathParams))
	if err != nil {
		return Action{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Action{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Action{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return Action{}, errors.New("could not read response body")
		}

		var response Action
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Action{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return Action{}, errors.New("the server returned an unknown status code")
	}
}

// Execute
func (client *ActionTag) Execute(actionId string, payload ActionExecuteRequest) (ActionExecuteResponse, error) {
	pathParams := make(map[string]interface{})
	pathParams["action_id"] = actionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/action/execute/:action_id", pathParams))
	if err != nil {
		return ActionExecuteResponse{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return ActionExecuteResponse{}, errors.New("could not marshal provided JSON data")
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", u.String(), reqBody)
	if err != nil {
		return ActionExecuteResponse{}, errors.New("could not create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return ActionExecuteResponse{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return ActionExecuteResponse{}, errors.New("could not read response body")
		}

		var response ActionExecuteResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ActionExecuteResponse{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return ActionExecuteResponse{}, errors.New("the server returned an unknown status code")
	}
}

// GetForm
func (client *ActionTag) GetForm() (FormContainer, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/action/form", pathParams))
	if err != nil {
		return FormContainer{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return FormContainer{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return FormContainer{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return FormContainer{}, errors.New("could not read response body")
		}

		var response FormContainer
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return FormContainer{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return FormContainer{}, errors.New("the server returned an unknown status code")
	}
}

// GetClasses
func (client *ActionTag) GetClasses() (ActionIndex, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/action/list", pathParams))
	if err != nil {
		return ActionIndex{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return ActionIndex{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return ActionIndex{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
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

	switch resp.StatusCode {
	default:
		return ActionIndex{}, errors.New("the server returned an unknown status code")
	}
}

// Create
func (client *ActionTag) Create(payload ActionCreate) (Message, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/action", pathParams))
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

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
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

	switch resp.StatusCode {
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// GetAll
func (client *ActionTag) GetAll() (ActionCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/action", pathParams))
	if err != nil {
		return ActionCollection{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return ActionCollection{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return ActionCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return ActionCollection{}, errors.New("could not read response body")
		}

		var response ActionCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ActionCollection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return ActionCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewActionTag(httpClient *http.Client, parser *sdkgen.Parser) *ActionTag {
	return &ActionTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
