
// BackendScopeTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "github.com/apioo/sdkgen-go/v2"
    "io"
    "net/http"
    "net/url"
    
)

type BackendScopeTag struct {
    internal *sdkgen.TagAbstract
}



// Create Creates a new scope
func (client *BackendScopeTag) Create(payload BackendScopeCreate) (*CommonMessage, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/scope", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Delete Deletes an existing scope
func (client *BackendScopeTag) Delete(scopeId string) (*CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["scope_id"] = scopeId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/scope/$scope_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("DELETE", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Get Returns a specific scope
func (client *BackendScopeTag) Get(scopeId string) (*BackendScope, error) {
    pathParams := make(map[string]interface{})
    pathParams["scope_id"] = scopeId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/scope/$scope_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendScope
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetAll Returns a paginated list of scopes
func (client *BackendScopeTag) GetAll(startIndex int, count int, search string) (*BackendScopeCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/scope", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendScopeCollection
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetCategories Returns all available scopes grouped by category
func (client *BackendScopeTag) GetCategories() (*BackendScopeCategories, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/scope/categories", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return nil, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendScopeCategories
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Update Updates an existing scope
func (client *BackendScopeTag) Update(scopeId string, payload BackendScopeUpdate) (*CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["scope_id"] = scopeId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/scope/$scope_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return nil, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return nil, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("PUT", u.String(), reqBody)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return &data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return nil, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return nil, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewBackendScopeTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendScopeTag {
	return &BackendScopeTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
