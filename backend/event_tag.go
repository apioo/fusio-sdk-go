// EventTag automatically generated by SDKgen please do not edit this file manually
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

type EventTag struct {
	internal *sdkgen.TagAbstract
}

// Delete
func (client *EventTag) Delete(eventId string) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["event_id"] = eventId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/event/$event_id<[0-9]+|^~>", pathParams))
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 410:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// Update
func (client *EventTag) Update(eventId string, payload EventUpdate) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["event_id"] = eventId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/event/$event_id<[0-9]+|^~>", pathParams))
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 410:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// Get
func (client *EventTag) Get(eventId string) (Event, error) {
	pathParams := make(map[string]interface{})
	pathParams["event_id"] = eventId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/event/$event_id<[0-9]+|^~>", pathParams))
	if err != nil {
		return Event{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return Event{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return Event{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Event{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Event
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Event{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Event{}, errors.New("could not unmarshal JSON response")
		}

		return Event{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Event{}, errors.New("could not unmarshal JSON response")
		}

		return Event{}, &MessageException{
			Payload: response,
		}
	case 410:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Event{}, errors.New("could not unmarshal JSON response")
		}

		return Event{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Event{}, errors.New("could not unmarshal JSON response")
		}

		return Event{}, &MessageException{
			Payload: response,
		}
	default:
		return Event{}, errors.New("the server returned an unknown status code")
	}
}

// Create
func (client *EventTag) Create(payload EventCreate) (Message, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/event", pathParams))
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// GetAll
func (client *EventTag) GetAll(startIndex int, count int, search string) (EventCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	u, err := url.Parse(client.internal.Parser.Url("/backend/event", pathParams))
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return EventCollection{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response EventCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return EventCollection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return EventCollection{}, errors.New("could not unmarshal JSON response")
		}

		return EventCollection{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return EventCollection{}, errors.New("could not unmarshal JSON response")
		}

		return EventCollection{}, &MessageException{
			Payload: response,
		}
	default:
		return EventCollection{}, errors.New("the server returned an unknown status code")
	}
}

// DeleteSubscription
func (client *EventTag) DeleteSubscription(subscriptionId string) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["subscription_id"] = subscriptionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/event/subscription/$subscription_id<[0-9]+>", pathParams))
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 410:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// UpdateSubscription
func (client *EventTag) UpdateSubscription(subscriptionId string, payload EventSubscriptionUpdate) (Message, error) {
	pathParams := make(map[string]interface{})
	pathParams["subscription_id"] = subscriptionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/event/subscription/$subscription_id<[0-9]+>", pathParams))
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 410:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// GetSubscription
func (client *EventTag) GetSubscription(subscriptionId string) (EventSubscription, error) {
	pathParams := make(map[string]interface{})
	pathParams["subscription_id"] = subscriptionId

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/event/subscription/$subscription_id<[0-9]+>", pathParams))
	if err != nil {
		return EventSubscription{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return EventSubscription{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return EventSubscription{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return EventSubscription{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response EventSubscription
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return EventSubscription{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return EventSubscription{}, errors.New("could not unmarshal JSON response")
		}

		return EventSubscription{}, &MessageException{
			Payload: response,
		}
	case 404:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return EventSubscription{}, errors.New("could not unmarshal JSON response")
		}

		return EventSubscription{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return EventSubscription{}, errors.New("could not unmarshal JSON response")
		}

		return EventSubscription{}, &MessageException{
			Payload: response,
		}
	default:
		return EventSubscription{}, errors.New("the server returned an unknown status code")
	}
}

// CreateSubscription
func (client *EventTag) CreateSubscription(payload EventSubscriptionCreate) (Message, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})

	u, err := url.Parse(client.internal.Parser.Url("/backend/event/subscription", pathParams))
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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Message{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 400:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return Message{}, errors.New("could not unmarshal JSON response")
		}

		return Message{}, &MessageException{
			Payload: response,
		}
	default:
		return Message{}, errors.New("the server returned an unknown status code")
	}
}

// GetAllSubscriptions
func (client *EventTag) GetAllSubscriptions(startIndex int, count int, search string) (EventSubscriptionCollection, error) {
	pathParams := make(map[string]interface{})

	queryParams := make(map[string]interface{})
	queryParams["startIndex"] = startIndex
	queryParams["count"] = count
	queryParams["search"] = search

	u, err := url.Parse(client.internal.Parser.Url("/backend/event/subscription", pathParams))
	if err != nil {
		return EventSubscriptionCollection{}, errors.New("could not parse url")
	}

	u.RawQuery = client.internal.Parser.Query(queryParams).Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return EventSubscriptionCollection{}, errors.New("could not create request")
	}

	resp, err := client.internal.HttpClient.Do(req)
	if err != nil {
		return EventSubscriptionCollection{}, errors.New("could not send request")
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return EventSubscriptionCollection{}, errors.New("could not read response body")
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var response EventSubscriptionCollection
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return EventSubscriptionCollection{}, errors.New("could not unmarshal JSON response")
		}

		return response, nil
	}

	switch resp.StatusCode {
	case 401:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return EventSubscriptionCollection{}, errors.New("could not unmarshal JSON response")
		}

		return EventSubscriptionCollection{}, &MessageException{
			Payload: response,
		}
	case 500:
		var response Message
		err = json.Unmarshal(respBody, &response)
		if err != nil {
			return EventSubscriptionCollection{}, errors.New("could not unmarshal JSON response")
		}

		return EventSubscriptionCollection{}, &MessageException{
			Payload: response,
		}
	default:
		return EventSubscriptionCollection{}, errors.New("the server returned an unknown status code")
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
