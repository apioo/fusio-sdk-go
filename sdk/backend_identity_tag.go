
// BackendIdentityTag automatically generated by SDKgen please do not edit this file manually
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

type BackendIdentityTag struct {
    internal *sdkgen.TagAbstract
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
        var response BackendIdentity
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendIdentity{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendIdentity{}, err
            }

            return BackendIdentity{}, &CommonMessageException{
                Payload: response,
            }
        case 404:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendIdentity{}, err
            }

            return BackendIdentity{}, &CommonMessageException{
                Payload: response,
            }
        case 410:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendIdentity{}, err
            }

            return BackendIdentity{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendIdentity{}, err
            }

            return BackendIdentity{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendIdentity{}, errors.New("the server returned an unknown status code")
    }
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
        var response CommonFormContainer
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return CommonFormContainer{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return CommonFormContainer{}, err
            }

            return CommonFormContainer{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return CommonFormContainer{}, err
            }

            return CommonFormContainer{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return CommonFormContainer{}, errors.New("the server returned an unknown status code")
    }
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
        var response BackendIdentityIndex
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendIdentityIndex{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendIdentityIndex{}, err
            }

            return BackendIdentityIndex{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendIdentityIndex{}, err
            }

            return BackendIdentityIndex{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendIdentityIndex{}, errors.New("the server returned an unknown status code")
    }
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
        var response BackendIdentityCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendIdentityCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendIdentityCollection{}, err
            }

            return BackendIdentityCollection{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendIdentityCollection{}, err
            }

            return BackendIdentityCollection{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendIdentityCollection{}, errors.New("the server returned an unknown status code")
    }
}



func NewBackendIdentityTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendIdentityTag {
	return &BackendIdentityTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
