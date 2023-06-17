// EventTag automatically generated by SDKgen please do not edit this file manually
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

type EventTag struct {
	internal *sdkgen.TagAbstract
}

// GetAll
func (client *EventTag) GetAll() (EventCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/consumer/event", pathParams))
	if err != nil {
		return EventCollection{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return EventCollection{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return EventCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return EventCollection{}, errors.New("could not read response body")
		}

		var response EventCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return EventCollection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return EventCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewEventTag(httpClient *http.Client, parser *sdkgen.Parser) *EventTag {
	return &EventTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
