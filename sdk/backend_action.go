
// backend_action automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

type BackendAction struct {
    Id int `json:"id"`
    Status int `json:"status"`
    Name string `json:"name"`
    Class string `json:"class"`
    Async bool `json:"async"`
    Config *BackendActionConfig `json:"config"`
    Metadata *CommonMetadata `json:"metadata"`
}

