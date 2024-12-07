
// BackendConnectionTag automatically generated by SDKgen please do not edit this file manually
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

type BackendConnectionTag struct {
    internal *sdkgen.TagAbstract
}



// Create 
func (client *BackendConnectionTag) Create(payload BackendConnectionCreate) (CommonMessage, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/connection", pathParams))
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
func (client *BackendConnectionTag) Delete(connectionId string) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["connection_id"] = connectionId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id<[0-9]+|^~>", pathParams))
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
func (client *BackendConnectionTag) Get(connectionId string) (BackendConnection, error) {
    pathParams := make(map[string]interface{})
    pathParams["connection_id"] = connectionId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return BackendConnection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendConnection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendConnection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendConnection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendConnection
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendConnection{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendConnection{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetAll 
func (client *BackendConnectionTag) GetAll(startIndex int, count int, search string) (BackendConnectionCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/connection", pathParams))
    if err != nil {
        return BackendConnectionCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendConnectionCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendConnectionCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendConnectionCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendConnectionCollection
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendConnectionCollection{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendConnectionCollection{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetClasses 
func (client *BackendConnectionTag) GetClasses() (BackendConnectionIndex, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/connection/list", pathParams))
    if err != nil {
        return BackendConnectionIndex{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendConnectionIndex{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendConnectionIndex{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendConnectionIndex{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendConnectionIndex
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendConnectionIndex{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendConnectionIndex{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetForm 
func (client *BackendConnectionTag) GetForm(class string) (CommonFormContainer, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["class"] = class

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/connection/form", pathParams))
    if err != nil {
        return CommonFormContainer{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return CommonFormContainer{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return CommonFormContainer{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return CommonFormContainer{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data CommonFormContainer
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return CommonFormContainer{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return CommonFormContainer{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetIntrospection 
func (client *BackendConnectionTag) GetIntrospection(connectionId string) (BackendConnectionIntrospectionEntities, error) {
    pathParams := make(map[string]interface{})
    pathParams["connection_id"] = connectionId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id<[0-9]+|^~>/introspection", pathParams))
    if err != nil {
        return BackendConnectionIntrospectionEntities{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendConnectionIntrospectionEntities{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendConnectionIntrospectionEntities{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendConnectionIntrospectionEntities{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendConnectionIntrospectionEntities
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendConnectionIntrospectionEntities{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendConnectionIntrospectionEntities{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetIntrospectionForEntity 
func (client *BackendConnectionTag) GetIntrospectionForEntity(connectionId string, entity string) (BackendConnectionIntrospectionEntity, error) {
    pathParams := make(map[string]interface{})
    pathParams["connection_id"] = connectionId
    pathParams["entity"] = entity

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id<[0-9]+|^~>/introspection/:entity", pathParams))
    if err != nil {
        return BackendConnectionIntrospectionEntity{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendConnectionIntrospectionEntity{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendConnectionIntrospectionEntity{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendConnectionIntrospectionEntity{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendConnectionIntrospectionEntity
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendConnectionIntrospectionEntity{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendConnectionIntrospectionEntity{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetRedirect 
func (client *BackendConnectionTag) GetRedirect(connectionId string) (BackendConnectionRedirectResponse, error) {
    pathParams := make(map[string]interface{})
    pathParams["connection_id"] = connectionId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id<[0-9]+|^~>/redirect", pathParams))
    if err != nil {
        return BackendConnectionRedirectResponse{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendConnectionRedirectResponse{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendConnectionRedirectResponse{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendConnectionRedirectResponse{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendConnectionRedirectResponse
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendConnectionRedirectResponse{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendConnectionRedirectResponse{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Update 
func (client *BackendConnectionTag) Update(connectionId string, payload BackendConnectionUpdate) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["connection_id"] = connectionId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/connection/$connection_id<[0-9]+|^~>", pathParams))
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




func NewBackendConnectionTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendConnectionTag {
	return &BackendConnectionTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
