// BackendPlanTag automatically generated by SDKgen please do not edit this file manually
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

type BackendPlanTag struct {
	internal *sdkgen.TagAbstract
}

// Delete
func (client *BackendPlanTag) Delete(planId string) (CommonMessage, error) {
	pathParams := make(map[string]interface{})
	pathParams["plan_id"] = planId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/plan/$plan_id<[0-9]+|^~>", pathParams))
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
func (client *BackendPlanTag) Update(planId string, payload BackendPlanUpdate) (CommonMessage, error) {
	pathParams := make(map[string]interface{})
	pathParams["plan_id"] = planId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/plan/$plan_id<[0-9]+|^~>", pathParams))
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
func (client *BackendPlanTag) Get(planId string) (BackendPlan, error) {
	pathParams := make(map[string]interface{})
	pathParams["plan_id"] = planId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/plan/$plan_id<[0-9]+|^~>", pathParams))
	if err != nil {
		return BackendPlan{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendPlan{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendPlan{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendPlan{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendPlan
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendPlan{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendPlan{}, err
		}

		return BackendPlan{}, &CommonMessageException{
			Payload: response,
		}
	case 404:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendPlan{}, err
		}

		return BackendPlan{}, &CommonMessageException{
			Payload: response,
		}
	case 410:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendPlan{}, err
		}

		return BackendPlan{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendPlan{}, err
		}

		return BackendPlan{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendPlan{}, errors.New("the server returned an unknown status code")
	}
}

// Create
func (client *BackendPlanTag) Create(payload BackendPlanCreate) (CommonMessage, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/plan", pathParams))
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
func (client *BackendPlanTag) GetAll(startIndex int, count int, search string) (BackendPlanCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	u, err := url.Parse(client.internal.Parser.Url("/backend/plan", pathParams))
	if err != nil {
		return BackendPlanCollection{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendPlanCollection{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendPlanCollection{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendPlanCollection{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendPlanCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendPlanCollection{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendPlanCollection{}, err
		}

		return BackendPlanCollection{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendPlanCollection{}, err
		}

		return BackendPlanCollection{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendPlanCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewBackendPlanTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendPlanTag {
	return &BackendPlanTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
