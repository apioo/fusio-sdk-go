// ConsumerPaymentTag automatically generated by SDKgen please do not edit this file manually
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

type ConsumerPaymentTag struct {
	internal *sdkgen.TagAbstract
}

// Checkout
func (client *ConsumerPaymentTag) Checkout(provider string, payload ConsumerPaymentCheckoutRequest) (ConsumerPaymentCheckoutResponse, error) {
	pathParams := make(map[string]interface{})
	pathParams["provider"] = provider

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/consumer/payment/:provider/checkout", pathParams))
	if err != nil {
		return ConsumerPaymentCheckoutResponse{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return ConsumerPaymentCheckoutResponse{}, err
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", u.String(), reqBody)
	if err != nil {
		return ConsumerPaymentCheckoutResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return ConsumerPaymentCheckoutResponse{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ConsumerPaymentCheckoutResponse{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response ConsumerPaymentCheckoutResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ConsumerPaymentCheckoutResponse{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ConsumerPaymentCheckoutResponse{}, err
		}

		return ConsumerPaymentCheckoutResponse{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ConsumerPaymentCheckoutResponse{}, err
		}

		return ConsumerPaymentCheckoutResponse{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return ConsumerPaymentCheckoutResponse{}, errors.New("the server returned an unknown status code")
	}
}

// Portal
func (client *ConsumerPaymentTag) Portal(provider string, payload ConsumerPaymentPortalRequest) (ConsumerPaymentPortalResponse, error) {
	pathParams := make(map[string]interface{})
	pathParams["provider"] = provider

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/consumer/payment/:provider/portal", pathParams))
	if err != nil {
		return ConsumerPaymentPortalResponse{}, err
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	raw, err := json.Marshal(payload)
	if err != nil {
		return ConsumerPaymentPortalResponse{}, err
	}

	var reqBody = bytes.NewReader(raw)

	req, err := http.NewRequest("POST", u.String(), reqBody)
	if err != nil {
		return ConsumerPaymentPortalResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return ConsumerPaymentPortalResponse{}, err
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ConsumerPaymentPortalResponse{}, err
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response ConsumerPaymentPortalResponse
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ConsumerPaymentPortalResponse{}, err
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ConsumerPaymentPortalResponse{}, err
		}

		return ConsumerPaymentPortalResponse{}, &CommonMessageException{
			Payload: response,
		}
	case 500:
		var response CommonMessage
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return ConsumerPaymentPortalResponse{}, err
		}

		return ConsumerPaymentPortalResponse{}, &CommonMessageException{
			Payload: response,
		}
	default:
		return ConsumerPaymentPortalResponse{}, errors.New("the server returned an unknown status code")
	}
}

func NewConsumerPaymentTag(httpClient *http.Client, parser *sdkgen.Parser) *ConsumerPaymentTag {
	return &ConsumerPaymentTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}