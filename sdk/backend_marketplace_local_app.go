
// backend_marketplace_local_app automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk
type BackendMarketplaceLocalApp struct {
    Version string `json:"version"`
    Description string `json:"description"`
    Screenshot string `json:"screenshot"`
    Website string `json:"website"`
    DownloadUrl string `json:"downloadUrl"`
    ShaHash string `json:"sha1Hash"`
    StartUrl string `json:"startUrl"`
    Remote BackendMarketplaceApp `json:"remote"`
}
