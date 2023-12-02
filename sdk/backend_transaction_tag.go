// BackendTransactionTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

import (
	"encoding/json"
	"errors"
	"github.com/apioo/sdkgen-go"
	"io"
	"net/http"
	"net/url"
)

type BackendTransactionTag struct {
	internal *sdkgen.TagAbstract
}

// Get
func (client *BackendTransactionTag) Get(transactionId string) (BackendTransaction, error) {
	pathParams := make(map[string]interface{})
	pathParams["transaction_id"] = transactionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/transaction/$transaction_id<[0-9]+>", pathParams))
	if err != nil {
		return BackendTransaction{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendTransaction{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendTransaction{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendTransaction{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendTransaction
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendTransaction{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendTransaction{}, err
		}

		return BackendTransaction{}, &CommonMessageException{
			Payload: response,
		}
	case 404:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendTransaction{}, err
		}

		return BackendTransaction{}, &CommonMessageException{
			Payload: response,
		}
	case 410:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendTransaction{}, err
		}

		return BackendTransaction{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendTransaction{}, err
		}

		return BackendTransaction{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendTransaction{}, errors.New("the server returned an unknown status code")
	}
}

// GetAll
func (client *BackendTransactionTag) GetAll(startIndex int, count int, search string, from string, to string, planId int, userId int, appId int, status string, provider string) (BackendTransactionCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search
	queryParams["from"] = from
	queryParams["to"] = to
	queryParams["planId"] = planId
	queryParams["userId"] = userId
	queryParams["appId"] = appId
	queryParams["status"] = status
	queryParams["provider"] = provider

	u, err := url.Parse(client.internal.Parser.Url("/backend/transaction", pathParams))
	if err != nil {
		return BackendTransactionCollection{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return BackendTransactionCollection{}, err
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return BackendTransactionCollection{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return BackendTransactionCollection{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response BackendTransactionCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendTransactionCollection{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendTransactionCollection{}, err
		}

		return BackendTransactionCollection{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return BackendTransactionCollection{}, err
		}

		return BackendTransactionCollection{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return BackendTransactionCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewBackendTransactionTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendTransactionTag {
	return &BackendTransactionTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
