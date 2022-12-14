// BackendStatisticTimeAverageResource automatically generated by SDKgen please do not edit this file manually
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

type BackendStatisticTimeAverageResource struct {
	url    string
	client *http.Client
}

// BackendActionStatisticGetTimeAverage
func (resource BackendStatisticTimeAverageResource) BackendActionStatisticGetTimeAverage(query BackendLogCollectionQuery) (StatisticChart, error) {
	url, err := url.Parse(resource.url)
	if err != nil {
		return StatisticChart{}, errors.New("could not parse url")
	}

	rawJson, err := json.Marshal(query)
	if err != nil {
		return StatisticChart{}, errors.New("could not marshall query")
	}

	parameters := url.Query()
	err = json.Unmarshal(rawJson, &parameters)
	url.RawQuery = parameters.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return StatisticChart{}, errors.New("could not create request")
	}

	resp, err := resource.client.Do(req)
	if err != nil {
		return StatisticChart{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return StatisticChart{}, errors.New("could not read response body")
	}

	var response StatisticChart

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return StatisticChart{}, errors.New("could not unmarshal JSON response")
	}

	return response, nil
}

func NewBackendStatisticTimeAverageResource(resource *sdkgen.Resource) *BackendStatisticTimeAverageResource {
	return &BackendStatisticTimeAverageResource{
		url:    resource.BaseUrl + "/backend/statistic/time_average",
		client: resource.HttpClient,
	}
}
