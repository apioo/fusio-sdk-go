// DashboardTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

import (
	"encoding/json"
	"errors"
	"github.com/apioo/sdkgen-go"
	"io"
	"net/http"
	"net/url"
)

type DashboardTag struct {
	internal *sdkgen.TagAbstract
}

// GetAll
func (client *DashboardTag) GetAll() (Dashboard, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/dashboard", pathParams))
	if err != nil {
		return Dashboard{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Dashboard{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Dashboard{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Dashboard{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Dashboard
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Dashboard{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Dashboard{}, errors.New("could not unmarshal JSON response")
		}

		return Dashboard{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Dashboard{}, errors.New("could not unmarshal JSON response")
		}

		return Dashboard{}, &MessageException{
			Payload: response,
		}
	default:
		return Dashboard{}, errors.New("the server returned an unknown status code")
	}
}

func NewDashboardTag(httpClient *http.Client, parser *sdkgen.Parser) *DashboardTag {
	return &DashboardTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
