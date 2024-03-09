
// BackendRateTag automatically generated by SDKgen please do not edit this file manually
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

type BackendRateTag struct {
    internal *sdkgen.TagAbstract
}



// Delete 
func (client *BackendRateTag) Delete(rateId string) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["rate_id"] = rateId

    queryParams := make(map[string]interface{})

    u, err := url.Parse(client.internal.Parser.Url("/backend/rate/$rate_id<[0-9]+|^~>", pathParams))
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
func (client *BackendRateTag) Update(rateId string, payload BackendRateUpdate) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["rate_id"] = rateId

    queryParams := make(map[string]interface{})

    u, err := url.Parse(client.internal.Parser.Url("/backend/rate/$rate_id<[0-9]+|^~>", pathParams))
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
func (client *BackendRateTag) Get(rateId string) (BackendRate, error) {
    pathParams := make(map[string]interface{})
    pathParams["rate_id"] = rateId

    queryParams := make(map[string]interface{})

    u, err := url.Parse(client.internal.Parser.Url("/backend/rate/$rate_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return BackendRate{}, err
    }

    u.RawQuery = client.internal.Parser.Query(queryParams).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendRate{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendRate{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendRate{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendRate
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendRate{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendRate{}, err
            }

            return BackendRate{}, &CommonMessageException{
                Payload: response,
            }
        case 404:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendRate{}, err
            }

            return BackendRate{}, &CommonMessageException{
                Payload: response,
            }
        case 410:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendRate{}, err
            }

            return BackendRate{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendRate{}, err
            }

            return BackendRate{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendRate{}, errors.New("the server returned an unknown status code")
    }
}

// Create 
func (client *BackendRateTag) Create(payload BackendRateCreate) (CommonMessage, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    u, err := url.Parse(client.internal.Parser.Url("/backend/rate", pathParams))
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
func (client *BackendRateTag) GetAll(startIndex int, count int, search string) (BackendRateCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search

    u, err := url.Parse(client.internal.Parser.Url("/backend/rate", pathParams))
    if err != nil {
        return BackendRateCollection{}, err
    }

    u.RawQuery = client.internal.Parser.Query(queryParams).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendRateCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendRateCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendRateCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendRateCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendRateCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendRateCollection{}, err
            }

            return BackendRateCollection{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendRateCollection{}, err
            }

            return BackendRateCollection{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendRateCollection{}, errors.New("the server returned an unknown status code")
    }
}



func NewBackendRateTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendRateTag {
	return &BackendRateTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
