
// backend_token_collection automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

type BackendTokenCollection struct {
    TotalResults int `json:"totalResults"`
    StartIndex int `json:"startIndex"`
    ItemsPerPage int `json:"itemsPerPage"`
    Entry []BackendToken `json:"entry"`
}

