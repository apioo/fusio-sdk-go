
// backend_event_update automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

type BackendEventUpdate struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Schema string `json:"schema"`
    Metadata *CommonMetadata `json:"metadata"`
}

