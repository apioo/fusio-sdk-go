
// backend_identity_collection automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

type BackendIdentityCollection struct {
    TotalResults int `json:"totalResults"`
    StartIndex int `json:"startIndex"`
    ItemsPerPage int `json:"itemsPerPage"`
    Entry []BackendIdentity `json:"entry"`
}

