// GrantTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package consumer

import (
	"github.com/apioo/sdkgen-go"
)

type GrantTag struct {
	internal *sdkgen.TagAbstract
}

// Delete
func (client *GrantTag) Delete(grantId string) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["grant_id"] = grantId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/consumer/grant/$grant_id&lt;[0-9]+&gt;", pathParams))
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

// GetAll
func (client *GrantTag) GetAll() (GrantCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/consumer/grant", pathParams))
	if err != nil {
		return GrantCollection{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return GrantCollection{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return GrantCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return GrantCollection{}, errors.New("could not read response body")
		}

		var response GrantCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return GrantCollection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	default:
		return GrantCollection{}, errors.New("the server returned an unknown status code")
	}
}

func NewGrantTag(httpClient *http.Client, parser *sdkgen.Parser) *GrantTag {
	return &GrantTag{
		internal: &sdkgen.TagAbstract{
			HttpClient: httpClient,
			Parser:     parser,
		},
	}
}
