
// ConsumerGrantTag automatically generated by SDKgen please do not edit this file manually
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

type ConsumerGrantTag struct {
    internal *sdkgen.TagAbstract
}



// Delete 
func (client *ConsumerGrantTag) Delete(grantId string) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["grant_id"] = grantId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/grant/$grant_id<[0-9]+>", pathParams))
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

// GetAll 
func (client *ConsumerGrantTag) GetAll(startIndex int, count int, search string) (ConsumerGrantCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/grant", pathParams))
    if err != nil {
        return ConsumerGrantCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return ConsumerGrantCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return ConsumerGrantCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return ConsumerGrantCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data ConsumerGrantCollection
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode >= 0 && statusCode <= 999 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerGrantCollection{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return ConsumerGrantCollection{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewConsumerGrantTag(httpClient *http.Client, parser *sdkgen.Parser) *ConsumerGrantTag {
	return &ConsumerGrantTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
