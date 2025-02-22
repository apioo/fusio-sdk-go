
// BackendTokenTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    
    "encoding/json"
    "errors"
    "fmt"
    "github.com/apioo/sdkgen-go/v2"
    "io"
    "net/http"
    "net/url"
    
)

type BackendTokenTag struct {
    internal *sdkgen.TagAbstract
}



// Get 
func (client *BackendTokenTag) Get(tokenId string) (*BackendToken, error) {
    pathParams := make(map[string]interface{})
    pathParams["token_id"] = tokenId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/token/$token_id<[0-9]+>", pathParams))
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
        var data BackendToken
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

// GetAll 
func (client *BackendTokenTag) GetAll(startIndex int, count int, search string, from string, to string, appId int, userId int, status int, scope string, ip string) (*BackendTokenCollection, error) {
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
        var data BackendTokenCollection
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




func NewBackendTokenTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendTokenTag {
	return &BackendTokenTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
