
// backend_schema_collection automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

// A paginated collection of schema objects
type BackendSchemaCollection struct {
    TotalResults int `json:"totalResults"`
    StartIndex int `json:"startIndex"`
    ItemsPerPage int `json:"itemsPerPage"`
    Entry []BackendSchema `json:"entry"`
}

