// backend_marketplace_remote_app automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package fusio

type BackendMarketplaceRemoteApp struct {
	Version     string                `json:"version"`
	Description string                `json:"description"`
	Screenshot  string                `json:"screenshot"`
	Website     string                `json:"website"`
	DownloadUrl string                `json:"downloadUrl"`
	ShaHash     string                `json:"sha1Hash"`
	StartUrl    string                `json:"startUrl"`
	Local       BackendMarketplaceApp `json:"local"`
}
