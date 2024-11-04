
// ConsumerPlanTag automatically generated by SDKgen please do not edit this file manually
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

type ConsumerPlanTag struct {
    internal *sdkgen.TagAbstract
}



// Get 
func (client *ConsumerPlanTag) Get(planId string) (ConsumerPlan, error) {
    pathParams := make(map[string]interface{})
    pathParams["plan_id"] = planId

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/plan/$plan_id<[0-9]+|^~>", pathParams))
    if err != nil {
        return ConsumerPlan{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return ConsumerPlan{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return ConsumerPlan{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return ConsumerPlan{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data ConsumerPlan
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 401 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerPlan{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerPlan{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 410 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerPlan{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerPlan{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return ConsumerPlan{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetAll 
func (client *ConsumerPlanTag) GetAll(startIndex int, count int, search string) (ConsumerPlanCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["count"] = count
    queryParams["search"] = search

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/consumer/plan", pathParams))
    if err != nil {
        return ConsumerPlanCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return ConsumerPlanCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return ConsumerPlanCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return ConsumerPlanCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data ConsumerPlanCollection
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 401 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerPlanCollection{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data CommonMessage
        err := json.Unmarshal(respBody, &data)

        return ConsumerPlanCollection{}, &CommonMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return ConsumerPlanCollection{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewConsumerPlanTag(httpClient *http.Client, parser *sdkgen.Parser) *ConsumerPlanTag {
	return &ConsumerPlanTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
