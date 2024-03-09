
// SystemConnectionTag automatically generated by SDKgen please do not edit this file manually
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

type SystemConnectionTag struct {
    internal *sdkgen.TagAbstract
}



// Callback 
func (client *SystemConnectionTag) Callback(name string) (CommonMessage, error) {
    pathParams := make(map[string]interface{})
    pathParams["name"] = name

    queryParams := make(map[string]interface{})

    u, err := url.Parse(client.internal.Parser.Url("/system/connection/:name/callback", pathParams))
    if err != nil {
        return CommonMessage{}, err
    }

    u.RawQuery = client.internal.Parser.Query(queryParams).Encode()


    req, err := http.NewRequest("GET", u.String(), nil)
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
        default:
            return CommonMessage{}, errors.New("the server returned an unknown status code")
    }
}



func NewSystemConnectionTag(httpClient *http.Client, parser *sdkgen.Parser) *SystemConnectionTag {
	return &SystemConnectionTag{
        internal: &sdkgen.TagAbstract{
            HttpClient: httpClient,
            Parser: parser,
        },
	}
}
