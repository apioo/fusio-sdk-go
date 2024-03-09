
// AuthorizationTag automatically generated by SDKgen please do not edit this file manually
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

type AuthorizationTag struct {
    internal *sdkgen.TagAbstract
}



// GetWhoami 
func (client *AuthorizationTag) GetWhoami() (BackendUser, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/authorization/whoami", pathParams))
    if err != nil {
        return BackendUser{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
        return BackendUser{}, err
    }


    resp, err := client.internal.HttpClient.Do(req)
    if err != nil {
        return BackendUser{}, err
    }

    defer resp.Body.Close()

    respBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return BackendUser{}, err
    }

    if resp.StatusCode >= 200 && resp.StatusCode < 300 {
        var response BackendUser
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return BackendUser{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return BackendUser{}, err
            }

            return BackendUser{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return BackendUser{}, errors.New("the server returned an unknown status code")
    }
}

// Revoke 
func (client *AuthorizationTag) Revoke() (CommonMessage, error) {
    pathParams := make(map[string]interface{})

    queryParams := make(map[string]interface{})

    var queryStructNames []string

    u, err := url.Parse(client.internal.Parser.Url("/authorization/revoke", pathParams))
    if err != nil {
        return CommonMessage{}, err
    }

    u.RawQuery = client.internal.Parser.QueryWithStruct(queryParams, queryStructNames).Encode()


    req, err := http.NewRequest("POST", u.String(), nil)
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
        var response CommonMessage
        err = json.Unmarshal(respBody, &response)
        if err != nil {
            return CommonMessage{}, err
        }

        return response, nil
    }

    switch resp.StatusCode {
        case 400:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return CommonMessage{}, err
            }

            return CommonMessage{}, &CommonMessageException{
                Payload: response,
            }
        case 500:
            var response CommonMessage
            err = json.Unmarshal(respBody, &response)
            if err != nil {
                return CommonMessage{}, err
            }

            return CommonMessage{}, &CommonMessageException{
                Payload: response,
            }
        default:
            return CommonMessage{}, errors.New("the server returned an unknown status code")
    }
}



func NewAuthorizationTag(httpClient *http.Client, parser *sdkgen.Parser) *AuthorizationTag {
	return &AuthorizationTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
