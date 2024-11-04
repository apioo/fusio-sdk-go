
// backend_operation automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

type BackendOperation struct {
    Id int `json:"id"`
    Status int `json:"status"`
    Active bool `json:"active"`
    Public bool `json:"public"`
    Stability int `json:"stability"`
    Description string `json:"description"`
    HttpMethod string `json:"httpMethod"`
    HttpPath string `json:"httpPath"`
    HttpCode int `json:"httpCode"`
    Name string `json:"name"`
    Parameters *BackendOperationParameters `json:"parameters"`
    Incoming string `json:"incoming"`
    Outgoing string `json:"outgoing"`
    Throws *BackendOperationThrows `json:"throws"`
    Action string `json:"action"`
    Costs int `json:"costs"`
    Scopes []string `json:"scopes"`
    Metadata *CommonMetadata `json:"metadata"`
}

