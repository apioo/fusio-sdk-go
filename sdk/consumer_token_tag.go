
// ConsumerTokenTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    
    "io"
    "net/http"
    "net/url"
    
)

type ConsumerTokenTag struct {
    internal *sdkgen.TagAbstract
}



// Delete 
func (client *ConsumerTokenTag) Delete(tokenId string) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["token_id"] = tokenId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/token/$token_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return CommonMessage{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


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
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 401 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return CommonMessage{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return CommonMessage{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 410 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return CommonMessage{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return CommonMessage{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return CommonMessage{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Update 
func (client *ConsumerTokenTag) Update(tokenId string, payload ConsumerTokenUpdate) (ConsumerTokenAccessToken, error) {
    pathParams := make(map[string]interface{})
    pathParams["token_id"] = tokenId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/token/$token_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return ConsumerTokenAccessToken{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return ConsumerTokenAccessToken{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("PUT", u.String(), reqBody)
    if err != nil {
        return ConsumerTokenAccessToken{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return ConsumerTokenAccessToken{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return ConsumerTokenAccessToken{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data ConsumerTokenAccessToken
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 400 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTokenAccessToken{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 401 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTokenAccessToken{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTokenAccessToken{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 410 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTokenAccessToken{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTokenAccessToken{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return ConsumerTokenAccessToken{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Get 
func (client *ConsumerTokenTag) Get(tokenId string) (ConsumerToken, error) {
    pathParams := make(map[string]interface{})
    pathParams["token_id"] = tokenId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/token/$token_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return ConsumerToken{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return ConsumerToken{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return ConsumerToken{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return ConsumerToken{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data ConsumerToken
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 401 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerToken{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerToken{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 410 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerToken{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerToken{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return ConsumerToken{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Create 
func (client *ConsumerTokenTag) Create(payload ConsumerTokenCreate) (ConsumerTokenAccessToken, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/token", pathParams))
    if err != nil {
        return ConsumerTokenAccessToken{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return ConsumerTokenAccessToken{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return ConsumerTokenAccessToken{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return ConsumerTokenAccessToken{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return ConsumerTokenAccessToken{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data ConsumerTokenAccessToken
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 400 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTokenAccessToken{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 401 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTokenAccessToken{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTokenAccessToken{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return ConsumerTokenAccessToken{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetAll 
func (client *ConsumerTokenTag) GetAll(startIndex int, count int, search string) (ConsumerTokenCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/token", pathParams))
    if err != nil {
        return ConsumerTokenCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return ConsumerTokenCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return ConsumerTokenCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return ConsumerTokenCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data ConsumerTokenCollection
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 401 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTokenCollection{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTokenCollection{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return ConsumerTokenCollection{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewConsumerTokenTag(httpClient *http.Client, parser *sdkgen.Parser) *ConsumerTokenTag {
	return &ConsumerTokenTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
