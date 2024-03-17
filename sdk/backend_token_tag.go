
// BackendTokenTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    
    "encoding/json"
    "errors"
    "github.com/apioo/sdkgen-go"
    "io"
    "net/http"
    "net/url"
    
)

type BackendTokenTag struct {
    internal *sdkgen.TagAbstract
}



// Get 
func (client *BackendTokenTag) Get(tokenId string) (BackendToken, error) {
    pathParams := make(map[string]interface{})
    pathParams["token_id"] = tokenId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/token/$token_id<[0-9]+>", pathParams))
    if err != nil {
        return BackendToken{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendToken{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendToken{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendToken{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendToken
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendToken{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendToken{}, err
            }

            return BackendToken{}, &CommonMessageException{
                Payload: response,
            }
        case 404:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendToken{}, err
            }

            return BackendToken{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendToken{}, err
            }

            return BackendToken{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendToken{}, errors.New("the server returned an unknown status code")
    }
}

// GetAll 
func (client *BackendTokenTag) GetAll(startIndex int, count int, search string, from string, to string, appId int, userId int, status int, scope string, ip string) (BackendTokenCollection, error) {
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

    u, err := url.Parse(client.internal.Parser.Url("/backend/token", pathParams))
    if err != nil {
        return BackendTokenCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendTokenCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendTokenCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendTokenCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendTokenCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendTokenCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendTokenCollection{}, err
            }

            return BackendTokenCollection{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendTokenCollection{}, err
            }

            return BackendTokenCollection{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendTokenCollection{}, errors.New("the server returned an unknown status code")
    }
}



func NewBackendTokenTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendTokenTag {
	return &BackendTokenTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}