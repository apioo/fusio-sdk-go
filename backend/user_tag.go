// UserTag automatically generated by SDKgen please do not edit this file manually
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

type UserTag struct {
	internal *sdkgen.TagAbstract
}

// Delete
func (client *UserTag) Delete(userId string) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["user_id"] = userId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/user/$user_id&lt;[0-9]+&gt;", pathParams))
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
func (client *UserTag) Update(userId string, payload UserUpdate) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["user_id"] = userId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/user/$user_id&lt;[0-9]+&gt;", pathParams))
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
func (client *UserTag) Get(userId string) (User, error) {
	pathParams := make(map[string]interface{})
	pathParams["user_id"] = userId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/user/$user_id&lt;[0-9]+&gt;", pathParams))
	if err != nil {
		return User{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return User{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return User{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return User{}, errors.New("could not read response body")
		}

		var response User
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return User{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return User{}, errors.New("the server returned an unknown status code")
	}
}

// Create
func (client *UserTag) Create(payload UserCreate) (Message, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/user", pathParams))
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
func (client *UserTag) GetAll() (UserCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/user", pathParams))
	if err != nil {
		return UserCollection{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return UserCollection{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return UserCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return UserCollection{}, errors.New("could not read response body")
		}

		var response UserCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return UserCollection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return UserCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewUserTag(httpClient *http.Client, parser *sdkgen.Parser) *UserTag {
	return &UserTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
