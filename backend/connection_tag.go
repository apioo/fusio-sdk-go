// ConnectionTag automatically generated by SDKgen please do not edit this file manually
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

type ConnectionTag struct {
	internal *sdkgen.TagAbstract
}

// GetIntrospectionForEntity
func (client *ConnectionTag) GetIntrospectionForEntity(connectionId string, entity string) (ConnectionIntrospectionEntity, error) {
	pathParams := make(map[string]interface{})
	pathParams["connection_id"] = connectionId
	pathParams["entity"] = entity

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id&lt;[0-9]+|^~&gt;/introspection/:entity", pathParams))
	if err != nil {
		return ConnectionIntrospectionEntity{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return ConnectionIntrospectionEntity{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return ConnectionIntrospectionEntity{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return ConnectionIntrospectionEntity{}, errors.New("could not read response body")
		}

		var response ConnectionIntrospectionEntity
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ConnectionIntrospectionEntity{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return ConnectionIntrospectionEntity{}, errors.New("the server returned an unknown status code")
	}
}

// GetIntrospection
func (client *ConnectionTag) GetIntrospection(connectionId string) (ConnectionIntrospectionEntities, error) {
	pathParams := make(map[string]interface{})
	pathParams["connection_id"] = connectionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id&lt;[0-9]+|^~&gt;/introspection", pathParams))
	if err != nil {
		return ConnectionIntrospectionEntities{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return ConnectionIntrospectionEntities{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return ConnectionIntrospectionEntities{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return ConnectionIntrospectionEntities{}, errors.New("could not read response body")
		}

		var response ConnectionIntrospectionEntities
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ConnectionIntrospectionEntities{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return ConnectionIntrospectionEntities{}, errors.New("the server returned an unknown status code")
	}
}

// GetRedirect
func (client *ConnectionTag) GetRedirect(connectionId string) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["connection_id"] = connectionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id&lt;[0-9]+|^~&gt;/redirect", pathParams))
	if err != nil {
		return Message{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
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

// Delete
func (client *ConnectionTag) Delete(connectionId string) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["connection_id"] = connectionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id&lt;[0-9]+|^~&gt;", pathParams))
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
func (client *ConnectionTag) Update(connectionId string, payload ConnectionUpdate) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["connection_id"] = connectionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id&lt;[0-9]+|^~&gt;", pathParams))
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
func (client *ConnectionTag) Get(connectionId string) (Connection, error) {
	pathParams := make(map[string]interface{})
	pathParams["connection_id"] = connectionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id&lt;[0-9]+|^~&gt;", pathParams))
	if err != nil {
		return Connection{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Connection{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Connection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return Connection{}, errors.New("could not read response body")
		}

		var response Connection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Connection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return Connection{}, errors.New("the server returned an unknown status code")
	}
}

// GetForm
func (client *ConnectionTag) GetForm() (FormContainer, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/form", pathParams))
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
func (client *ConnectionTag) GetClasses() (ConnectionIndex, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection/list", pathParams))
	if err != nil {
		return ConnectionIndex{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return ConnectionIndex{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return ConnectionIndex{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
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

	switch resp.StatusCode {
	default:
		return ConnectionIndex{}, errors.New("the server returned an unknown status code")
	}
}

// Create
func (client *ConnectionTag) Create(payload ConnectionCreate) (ConnectionCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection", pathParams))
	if err != nil {
		return ConnectionCollection{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return ConnectionCollection{}, errors.New("could not marshal provided JSON data")
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", u.String(), reqBody)
	if err != nil {
		return ConnectionCollection{}, errors.New("could not create request")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return ConnectionCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return ConnectionCollection{}, errors.New("could not read response body")
		}

		var response ConnectionCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ConnectionCollection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return ConnectionCollection{}, errors.New("the server returned an unknown status code")
	}
}

// GetAll
func (client *ConnectionTag) GetAll() (ConnectionCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/connection", pathParams))
	if err != nil {
		return ConnectionCollection{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return ConnectionCollection{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return ConnectionCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return ConnectionCollection{}, errors.New("could not read response body")
		}

		var response ConnectionCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ConnectionCollection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return ConnectionCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewConnectionTag(httpClient *http.Client, parser *sdkgen.Parser) *ConnectionTag {
	return &ConnectionTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
