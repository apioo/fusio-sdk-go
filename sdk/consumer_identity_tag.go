
// ConsumerIdentityTag automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import (
    
    "encoding/json"
    "errors"
    "fmt"
    
    "io"
    "net/http"
    "net/url"
    
)

type ConsumerIdentityTag struct {
    internal *sdkgen.TagAbstract
}



// Exchange 
func (client *ConsumerIdentityTag) Exchange(identity string) (Passthru, error) {
    pathParams := make(map[string]interface{})
    pathParams["identity"] = identity

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/identity/:identity/exchange", pathParams))
    if err != nil {
        return Passthru{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return Passthru{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return Passthru{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return Passthru{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data Passthru
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return Passthru{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return Passthru{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetAll 
func (client *ConsumerIdentityTag) GetAll(appId int, appKey string) (ConsumerIdentityCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["appId"] = appId
    queryParams["appKey"] = appKey

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/identity", pathParams))
    if err != nil {
        return ConsumerIdentityCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return ConsumerIdentityCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return ConsumerIdentityCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return ConsumerIdentityCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data ConsumerIdentityCollection
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerIdentityCollection{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return ConsumerIdentityCollection{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Redirect 
func (client *ConsumerIdentityTag) Redirect(identity string) (Passthru, error) {
    pathParams := make(map[string]interface{})
    pathParams["identity"] = identity

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/identity/:identity/redirect", pathParams))
    if err != nil {
        return Passthru{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return Passthru{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return Passthru{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return Passthru{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data Passthru
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return Passthru{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return Passthru{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewConsumerIdentityTag(httpClient *http.Client, parser *sdkgen.Parser) *ConsumerIdentityTag {
	return &ConsumerIdentityTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
