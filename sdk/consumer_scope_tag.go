
// ConsumerScopeTag automatically generated by SDKgen please do not edit this file manually
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

type ConsumerScopeTag struct {
    internal *sdkgen.TagAbstract
}



// GetAll 
func (client *ConsumerScopeTag) GetAll(startIndex int, count int, search string) (ConsumerScopeCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/scope", pathParams))
    if err != nil {
        return ConsumerScopeCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return ConsumerScopeCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return ConsumerScopeCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return ConsumerScopeCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response ConsumerScopeCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return ConsumerScopeCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return ConsumerScopeCollection{}, err
            }

            return ConsumerScopeCollection{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return ConsumerScopeCollection{}, err
            }

            return ConsumerScopeCollection{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return ConsumerScopeCollection{}, errors.New("the server returned an unknown status code")
    }
}



func NewConsumerScopeTag(httpClient *http.Client, parser *sdkgen.Parser) *ConsumerScopeTag {
	return &ConsumerScopeTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
