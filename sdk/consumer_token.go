
// consumer_token automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import "time"
type ConsumerToken struct {
    Id int `json:"id"`
    Status int `json:"status"`
    Name string `json:"name"`
    Scopes []string `json:"scopes"`
    Ip string `json:"ip"`
    Expire time.Time `json:"expire"`
    Date time.Time `json:"date"`
}
