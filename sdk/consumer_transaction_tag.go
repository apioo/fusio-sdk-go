
// ConsumerTransactionTag automatically generated by SDKgen please do not edit this file manually
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

type ConsumerTransactionTag struct {
    internal *sdkgen.TagAbstract
}



// Get 
func (client *ConsumerTransactionTag) Get(transactionId string) (ConsumerTransaction, error) {
    pathParams := make(map[string]interface{})
    pathParams["transaction_id"] = transactionId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/transaction/$transaction_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return ConsumerTransaction{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return ConsumerTransaction{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return ConsumerTransaction{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return ConsumerTransaction{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data ConsumerTransaction
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 401 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTransaction{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTransaction{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 410 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTransaction{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTransaction{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return ConsumerTransaction{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetAll 
func (client *ConsumerTransactionTag) GetAll(startIndex int, count int, search string) (ConsumerTransactionCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/transaction", pathParams))
    if err != nil {
        return ConsumerTransactionCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return ConsumerTransactionCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return ConsumerTransactionCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return ConsumerTransactionCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data ConsumerTransactionCollection
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 401 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTransactionCollection{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerTransactionCollection{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return ConsumerTransactionCollection{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewConsumerTransactionTag(httpClient *http.Client, parser *sdkgen.Parser) *ConsumerTransactionTag {
	return &ConsumerTransactionTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
