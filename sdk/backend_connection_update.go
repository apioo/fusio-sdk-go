
// backend_connection_update automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

type BackendConnectionUpdate struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Class string `json:"class"`
    Oauth bool `json:"oauth2"`
    Config *BackendConnectionConfig `json:"config"`
    Metadata *CommonMetadata `json:"metadata"`
}

