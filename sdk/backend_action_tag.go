
// BackendActionTag automatically generated by SDKgen please do not edit this file manually
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

type BackendActionTag struct {
    internal *sdkgen.TagAbstract
}



// Delete 
func (client *BackendActionTag) Delete(actionId string) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["action_id"] = actionId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/action/$action_id<[0-9]+|^~>", pathParams))
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
        case 404:
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
func (client *BackendActionTag) Update(actionId string, payload BackendActionUpdate) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["action_id"] = actionId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/action/$action_id<[0-9]+|^~>", pathParams))
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
func (client *BackendActionTag) Get(actionId string) (BackendAction, error) {
    pathParams := make(map[string]interface{})
    pathParams["action_id"] = actionId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/action/$action_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return BackendAction{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendAction{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendAction{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendAction{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendAction
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendAction{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendAction{}, err
            }

            return BackendAction{}, &CommonMessageException{
                Payload: response,
            }
        case 404:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendAction{}, err
            }

            return BackendAction{}, &CommonMessageException{
                Payload: response,
            }
        case 410:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendAction{}, err
            }

            return BackendAction{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendAction{}, err
            }

            return BackendAction{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendAction{}, errors.New("the server returned an unknown status code")
    }
}

// Execute 
func (client *BackendActionTag) Execute(actionId string, payload BackendActionExecuteRequest) (BackendActionExecuteResponse, error) {
    pathParams := make(map[string]interface{})
    pathParams["action_id"] = actionId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/action/execute/:action_id", pathParams))
    if err != nil {
        return BackendActionExecuteResponse{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return BackendActionExecuteResponse{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return BackendActionExecuteResponse{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendActionExecuteResponse{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendActionExecuteResponse{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendActionExecuteResponse
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendActionExecuteResponse{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendActionExecuteResponse{}, err
            }

            return BackendActionExecuteResponse{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendActionExecuteResponse{}, err
            }

            return BackendActionExecuteResponse{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendActionExecuteResponse{}, errors.New("the server returned an unknown status code")
    }
}

// GetForm 
func (client *BackendActionTag) GetForm(class string) (CommonFormContainer, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["class"] = class

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/action/form", pathParams))
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
func (client *BackendActionTag) GetClasses() (BackendActionIndex, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/action/list", pathParams))
    if err != nil {
        return BackendActionIndex{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendActionIndex{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendActionIndex{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendActionIndex{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendActionIndex
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendActionIndex{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendActionIndex{}, err
            }

            return BackendActionIndex{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendActionIndex{}, err
            }

            return BackendActionIndex{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendActionIndex{}, errors.New("the server returned an unknown status code")
    }
}

// Create 
func (client *BackendActionTag) Create(payload BackendActionCreate) (CommonMessage, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/action", pathParams))
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
func (client *BackendActionTag) GetAll(startIndex int, count int, search string) (BackendActionCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/action", pathParams))
    if err != nil {
        return BackendActionCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendActionCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendActionCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendActionCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendActionCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendActionCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendActionCollection{}, err
            }

            return BackendActionCollection{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendActionCollection{}, err
            }

            return BackendActionCollection{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendActionCollection{}, errors.New("the server returned an unknown status code")
    }
}



func NewBackendActionTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendActionTag {
	return &BackendActionTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
