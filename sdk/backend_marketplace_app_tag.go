
// BackendMarketplaceAppTag automatically generated by SDKgen please do not edit this file manually
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

type BackendMarketplaceAppTag struct {
    internal *sdkgen.TagAbstract
}



// Upgrade 
func (client *BackendMarketplaceAppTag) Upgrade(user string, name string) (MarketplaceMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["user"] = user
    pathParams["name"] = name

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/marketplace/app/:user/:name", pathParams))
    if err != nil {
        return MarketplaceMessage{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("PUT", u.String(), nil)
    if err != nil {
        return MarketplaceMessage{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return MarketplaceMessage{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return MarketplaceMessage{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 400 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceMessage{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 401 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceMessage{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceMessage{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 410 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceMessage{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceMessage{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return MarketplaceMessage{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Get 
func (client *BackendMarketplaceAppTag) Get(user string, name string) (MarketplaceApp, error) {
    pathParams := make(map[string]interface{})
    pathParams["user"] = user
    pathParams["name"] = name

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/marketplace/app/:user/:name", pathParams))
    if err != nil {
        return MarketplaceApp{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return MarketplaceApp{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return MarketplaceApp{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return MarketplaceApp{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data MarketplaceApp
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 401 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceApp{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 404 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceApp{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 410 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceApp{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceApp{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return MarketplaceApp{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// Install 
func (client *BackendMarketplaceAppTag) Install(payload MarketplaceInstall) (MarketplaceMessage, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/marketplace/app", pathParams))
    if err != nil {
        return MarketplaceMessage{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()

    raw, err := json.Marshal(payload)
    if err != nil {
        return MarketplaceMessage{}, err
    }

    var reqBody = bytes.NewReader(raw)

    req, err := http.NewRequest("POST", u.String(), reqBody)
    if err != nil {
        return MarketplaceMessage{}, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return MarketplaceMessage{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return MarketplaceMessage{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 400 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceMessage{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 401 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceMessage{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceMessage{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return MarketplaceMessage{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}

// GetAll 
func (client *BackendMarketplaceAppTag) GetAll(startIndex int, query string) (MarketplaceAppCollection, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})
    queryParams["startIndex"] = startIndex
    queryParams["query"] = query

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/backend/marketplace/app", pathParams))
    if err != nil {
        return MarketplaceAppCollection{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return MarketplaceAppCollection{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return MarketplaceAppCollection{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return MarketplaceAppCollection{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var data MarketplaceAppCollection
        err := json.Unmarshal(respBody, &data)

        return data, err
    }

    var statusCode = resp.StatusCode
    if statusCode == 401 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceAppCollection{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    if statusCode == 500 {
        var data MarketplaceMessage
        err := json.Unmarshal(respBody, &data)

        return MarketplaceAppCollection{}, &MarketplaceMessageException{
            Payload: data,
            Previous: err,
        }
    }

    return MarketplaceAppCollection{}, errors.New(fmt.Sprint("The server returned an unknown status code: ", statusCode))
}




func NewBackendMarketplaceAppTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendMarketplaceAppTag {
	return &BackendMarketplaceAppTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
