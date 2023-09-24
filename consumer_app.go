// consumer_app automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package fusio

type ConsumerApp struct {
	Id        int            `json:"id"`
	UserId    int            `json:"userId"`
	Status    int            `json:"status"`
	Name      string         `json:"name"`
	Url       string         `json:"url"`
	AppKey    string         `json:"appKey"`
	AppSecret string         `json:"appSecret"`
	Date      string         `json:"date"`
	Scopes    []string       `json:"scopes"`
	Metadata  CommonMetadata `json:"metadata"`
}
