// BackendSdkTag automatically generated by SDKgen please do not edit this file manually
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

type BackendSdkTag struct {
	internal *sdkgen.TagAbstract
}

// Generate
func (client *BackendSdkTag) Generate(payload BackendSdkGenerate) (CommonMessage, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/sdk", pathParams))
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
func (client *BackendSdkTag) GetAll() (BackendSdkResponse, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/sdk", pathParams))
	if err != nil {
		return BackendSdkResponse{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendSdkResponse{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendSdkResponse{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendSdkResponse{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendSdkResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendSdkResponse{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendSdkResponse{}, err
		}

		return BackendSdkResponse{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendSdkResponse{}, err
		}

		return BackendSdkResponse{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendSdkResponse{}, errors.New("the server returned an unknown status code")
	}
}

func NewBackendSdkTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendSdkTag {
	return &BackendSdkTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
