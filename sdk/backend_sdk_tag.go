
// BackendSdkTag automatically generated by SDKgen please do not edit this file manually
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

type BackendSdkTag struct {
    internal *sdkgen.TagAbstract
}



// Generate 
func (client *BackendSdkTag) Generate(payload BackendSdkGenerate) (BackendSdkMessage, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/sdk", pathParams))
    if err != nil {
        return BackendSdkMessage{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return BackendSdkMessage{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return BackendSdkMessage{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendSdkMessage{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendSdkMessage{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendSdkMessage
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendSdkMessage{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendSdkMessage{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetAll 
func (client *BackendSdkTag) GetAll() (BackendSdkResponse, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/sdk", pathParams))
    if err != nil {
        return BackendSdkResponse{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendSdkResponse{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendSdkResponse{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendSdkResponse{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data BackendSdkResponse
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return BackendSdkResponse{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return BackendSdkResponse{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewBackendSdkTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendSdkTag {
	return &BackendSdkTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
