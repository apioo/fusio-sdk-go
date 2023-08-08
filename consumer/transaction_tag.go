// TransactionTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package consumer

import (
	"encoding/json"
	"errors"
	"github.com/apioo/sdkgen-go"
	"io"
	"net/http"
	"net/url"
)

type TransactionTag struct {
	internal *sdkgen.TagAbstract
}

// Get
func (client *TransactionTag) Get(transactionId string) (Transaction, error) {
	pathParams := make(map[string]interface{})
	pathParams["transaction_id"] = transactionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/consumer/transaction/$transaction_id<[0-9]+>", pathParams))
	if err != nil {
		return Transaction{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Transaction{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Transaction{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Transaction{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Transaction
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Transaction{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Transaction{}, errors.New("could not unmarshal JSON response")
		}

		return Transaction{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Transaction{}, errors.New("could not unmarshal JSON response")
		}

		return Transaction{}, &MessageException{
			Payload: response,
		}
	case 410:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Transaction{}, errors.New("could not unmarshal JSON response")
		}

		return Transaction{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Transaction{}, errors.New("could not unmarshal JSON response")
		}

		return Transaction{}, &MessageException{
			Payload: response,
		}
	default:
		return Transaction{}, errors.New("the server returned an unknown status code")
	}
}

// GetAll
func (client *TransactionTag) GetAll(startIndex int, count int, search string) (TransactionCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	u, err := url.Parse(client.internal.Parser.Url("/consumer/transaction", pathParams))
	if err != nil {
		return TransactionCollection{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return TransactionCollection{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return TransactionCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return TransactionCollection{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response TransactionCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TransactionCollection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TransactionCollection{}, errors.New("could not unmarshal JSON response")
		}

		return TransactionCollection{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return TransactionCollection{}, errors.New("could not unmarshal JSON response")
		}

		return TransactionCollection{}, &MessageException{
			Payload: response,
		}
	default:
		return TransactionCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewTransactionTag(httpClient *http.Client, parser *sdkgen.Parser) *TransactionTag {
	return &TransactionTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}