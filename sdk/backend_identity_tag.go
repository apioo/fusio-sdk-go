
// BackendIdentityTag automatically generated by SDKgen please do not edit this file manually
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

type BackendIdentityTag struct {
    internal *sdkgen.TagAbstract
}



// Create 
func (client *BackendIdentityTag) Create(payload BackendIdentityCreate) (CommonMessage, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/identity", pathParams))
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
func (client *BackendIdentityTag) Delete(identityId string) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["identity_id"] = identityId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/identity/$identity_id<[0-9]+|^~>", pathParams))
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
func (client *BackendIdentityTag) Get(identityId string) (BackendIdentity, error) {
    pathParams := make(map[string]interface{})
    pathParams["identity_id"] = identityId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/identity/$identity_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return BackendIdentity{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendIdentity{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendIdentity{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendIdentity{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendIdentity
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendIdentity{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendIdentity{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetAll 
func (client *BackendIdentityTag) GetAll(startIndex int, count int, search string) (BackendIdentityCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/identity", pathParams))
    if err != nil {
        return BackendIdentityCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendIdentityCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendIdentityCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendIdentityCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendIdentityCollection
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendIdentityCollection{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendIdentityCollection{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetClasses 
func (client *BackendIdentityTag) GetClasses() (BackendIdentityIndex, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/identity/list", pathParams))
    if err != nil {
        return BackendIdentityIndex{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendIdentityIndex{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendIdentityIndex{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendIdentityIndex{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendIdentityIndex
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendIdentityIndex{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendIdentityIndex{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetForm 
func (client *BackendIdentityTag) GetForm(class string) (CommonFormContainer, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["class"] = class

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/identity/form", pathParams))
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

// Update 
func (client *BackendIdentityTag) Update(identityId string, payload BackendIdentityUpdate) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["identity_id"] = identityId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/identity/$identity_id<[0-9]+|^~>", pathParams))
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




func NewBackendIdentityTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendIdentityTag {
	return &BackendIdentityTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
