
// backend_connection_create automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk
type BackendConnectionCreate struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Class string `json:"class"`
    Config *BackendConnectionConfig `json:"config"`
    Metadata *CommonMetadata `json:"metadata"`
}
