
// backend_scope_create automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

type BackendScopeCreate struct {
    Id int `json:"id"`
    Name string `json:"name"`
    Description string `json:"description"`
    Operations []BackendScopeOperation `json:"operations"`
    Metadata *CommonMetadata `json:"metadata"`
}

