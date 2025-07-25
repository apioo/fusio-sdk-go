
// backend_event automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

// This object represents an event which can be triggered by an action
type BackendEvent struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Schema string `json:"schema"`
    Metadata *CommonMetadata `json:"metadata"`
}

