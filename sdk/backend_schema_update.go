
// backend_schema_update automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

type BackendSchemaUpdate struct {
    Id int `json:"id"`
    Status int `json:"status"`
    Name string `json:"name"`
    Source *BackendSchemaSource `json:"source"`
    Form *BackendSchemaForm `json:"form"`
    Readonly bool `json:"readonly"`
    Metadata *CommonMetadata `json:"metadata"`
}

