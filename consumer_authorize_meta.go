// consumer_authorize_meta automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package fusio

type ConsumerAuthorizeMeta struct {
	Name   string          `json:"name"`
	Url    string          `json:"url"`
	Scopes []ConsumerScope `json:"scopes"`
}
