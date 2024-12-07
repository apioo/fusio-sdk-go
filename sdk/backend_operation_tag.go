
// BackendOperationTag automatically generated by SDKgen please do not edit this file manually
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

type BackendOperationTag struct {
    internal *sdkgen.TagAbstract
}



// Create 
func (client *BackendOperationTag) Create(payload BackendOperationCreate) (CommonMessage, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/operation", pathParams))
    if err != nil {
        return CommonMessage{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

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
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return CommonMessage{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return CommonMessage{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Delete 
func (client *BackendOperationTag) Delete(operationId string) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["operation_id"] = operationId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/operation/$operation_id<[0-9]+|^~>", pathParams))
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
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return CommonMessage{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return CommonMessage{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Get 
func (client *BackendOperationTag) Get(operationId string) (BackendOperation, error) {
    pathParams := make(map[string]interface{})
    pathParams["operation_id"] = operationId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/operation/$operation_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return BackendOperation{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendOperation{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendOperation{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendOperation{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendOperation
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendOperation{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendOperation{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetAll 
func (client *BackendOperationTag) GetAll(startIndex int, count int, search string) (BackendOperationCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/operation", pathParams))
    if err != nil {
        return BackendOperationCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendOperationCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendOperationCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendOperationCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendOperationCollection
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendOperationCollection{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendOperationCollection{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Update 
func (client *BackendOperationTag) Update(operationId string, payload BackendOperationUpdate) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["operation_id"] = operationId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/operation/$operation_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return CommonMessage{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

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
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return CommonMessage{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return CommonMessage{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewBackendOperationTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendOperationTag {
	return &BackendOperationTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
