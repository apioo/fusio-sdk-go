
// consumer_token_create automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import "time"
type ConsumerTokenCreate struct {
    Name string `json:"name"`
    Scope []string `json:"scope"`
    Expire time.Time `json:"expire"`
}