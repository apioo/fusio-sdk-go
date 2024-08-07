
// BackendMarketplaceAppTag automatically generated by SDKgen please do not edit this file manually
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
        var response MarketplaceMessage
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return MarketplaceMessage{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 400:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceMessage{}, err
            }

            return MarketplaceMessage{}, &MarketplaceMessageException{
                Payload: response,
            }
        case 401:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceMessage{}, err
            }

            return MarketplaceMessage{}, &MarketplaceMessageException{
                Payload: response,
            }
        case 404:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceMessage{}, err
            }

            return MarketplaceMessage{}, &MarketplaceMessageException{
                Payload: response,
            }
        case 410:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceMessage{}, err
            }

            return MarketplaceMessage{}, &MarketplaceMessageException{
                Payload: response,
            }
        case 500:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceMessage{}, err
            }

            return MarketplaceMessage{}, &MarketplaceMessageException{
                Payload: response,
            }
        default:
            return MarketplaceMessage{}, errors.New("the server returned an unknown status code")
    }
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
        var response MarketplaceApp
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return MarketplaceApp{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceApp{}, err
            }

            return MarketplaceApp{}, &MarketplaceMessageException{
                Payload: response,
            }
        case 404:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceApp{}, err
            }

            return MarketplaceApp{}, &MarketplaceMessageException{
                Payload: response,
            }
        case 410:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceApp{}, err
            }

            return MarketplaceApp{}, &MarketplaceMessageException{
                Payload: response,
            }
        case 500:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceApp{}, err
            }

            return MarketplaceApp{}, &MarketplaceMessageException{
                Payload: response,
            }
        default:
            return MarketplaceApp{}, errors.New("the server returned an unknown status code")
    }
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
        var response MarketplaceMessage
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return MarketplaceMessage{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 400:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceMessage{}, err
            }

            return MarketplaceMessage{}, &MarketplaceMessageException{
                Payload: response,
            }
        case 401:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceMessage{}, err
            }

            return MarketplaceMessage{}, &MarketplaceMessageException{
                Payload: response,
            }
        case 500:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceMessage{}, err
            }

            return MarketplaceMessage{}, &MarketplaceMessageException{
                Payload: response,
            }
        default:
            return MarketplaceMessage{}, errors.New("the server returned an unknown status code")
    }
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
        var response MarketplaceAppCollection
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return MarketplaceAppCollection{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 401:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceAppCollection{}, err
            }

            return MarketplaceAppCollection{}, &MarketplaceMessageException{
                Payload: response,
            }
        case 500:
            var response MarketplaceMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return MarketplaceAppCollection{}, err
            }

            return MarketplaceAppCollection{}, &MarketplaceMessageException{
                Payload: response,
            }
        default:
            return MarketplaceAppCollection{}, errors.New("the server returned an unknown status code")
    }
}



func NewBackendMarketplaceAppTag(httpClient *http.Client, parser *sdkgen.Parser) *BackendMarketplaceAppTag {
	return &BackendMarketplaceAppTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
