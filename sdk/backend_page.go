
// backend_page automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

// This object represents a page at the developer portal
type BackendPage struct {
    Id int `json:"id"`
    Status int `json:"status"`
    Title string `json:"title"`
    Slug string `json:"slug"`
    Content string `json:"content"`
    Metadata *CommonMetadata `json:"metadata"`
}

