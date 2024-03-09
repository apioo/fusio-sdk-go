
// BackendAppTag automatically generated by SDKgen please do not edit this file manually
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

type BackendAppTag struct {
    internal *sdkgen.TagAbstract
}



// DeleteToken 
func (client *BackendAppTag) DeleteToken(appId string, tokenId string) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["app_id"] = appId
    pathParams["token_id"] = tokenId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/app/$app_id<[0-9]+>/token/:token_id", pathParams))
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

// Delete 
func (client *BackendAppTag) Delete(appId string) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["app_id"] = appId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/app/$app_id<[0-9]+>", pathParams))
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
func (client *BackendAppTag) Update(appId string, payload BackendAppUpdate) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["app_id"] = appId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/app/$app_id<[0-9]+>", pathParams))
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
func (client *BackendAppTag) Get(appId string) (BackendApp, error) {
    pathParams := make(map[string]interface{})
    pathParams["app_id"] = appId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/app/$app_id<[0-9]+>", pathParams))
    if err != nil {
        return BackendApp{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendApp{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendApp{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendApp{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendApp
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendApp{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 404:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendApp{}, err
            }

            return BackendApp{}, &CommonMessageException{
                Payload: response,
            }
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendApp{}, err
            }

            return BackendApp{}, &CommonMessageException{
                Payload: response,
            }
        case 410:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendApp{}, err
            }

            return BackendApp{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendApp{}, err
            }

            return BackendApp{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendApp{}, errors.New("the server returned an unknown status code")
    }
}

// Create 
func (client *BackendAppTag) Create(payload BackendAppCreate) (CommonMessage, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/app", pathParams))
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
func (client *BackendAppTag) GetAll(startIndex int, count int, search string) (BackendAppCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/app", pathParams))
    if err != nil {
        return BackendAppCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendAppCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendAppCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendAppCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendAppCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendAppCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendAppCollection{}, err
            }

            return BackendAppCollection{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendAppCollection{}, err
            }

            return BackendAppCollection{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendAppCollection{}, errors.New("the server returned an unknown status code")
    }
}

// GetToken 
func (client *BackendAppTag) GetToken(tokenId string) (BackendAppToken, error) {
    pathParams := make(map[string]interface{})
    pathParams["token_id"] = tokenId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/app/token/$token_id<[0-9]+>", pathParams))
    if err != nil {
        return BackendAppToken{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendAppToken{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendAppToken{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendAppToken{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendAppToken
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendAppToken{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendAppToken{}, err
            }

            return BackendAppToken{}, &CommonMessageException{
                Payload: response,
            }
        case 404:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendAppToken{}, err
            }

            return BackendAppToken{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendAppToken{}, err
            }

            return BackendAppToken{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendAppToken{}, errors.New("the server returned an unknown status code")
    }
}

// GetAllTokens 
func (client *BackendAppTag) GetAllTokens(startIndex int, count int, search string, from string, to string, appId int, userId int, status int, scope string, ip string) (BackendAppTokenCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search
    queryParams["from"] = from
    queryParams["to"] = to
    queryParams["appId"] = appId
    queryParams["userId"] = userId
    queryParams["status"] = status
    queryParams["scope"] = scope
    queryParams["ip"] = ip

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/app/token", pathParams))
    if err != nil {
        return BackendAppTokenCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendAppTokenCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendAppTokenCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendAppTokenCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendAppTokenCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendAppTokenCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendAppTokenCollection{}, err
            }

            return BackendAppTokenCollection{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendAppTokenCollection{}, err
            }

            return BackendAppTokenCollection{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendAppTokenCollection{}, errors.New("the server returned an unknown status code")
    }
}



func NewBackendAppTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendAppTag {
	return &BackendAppTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
