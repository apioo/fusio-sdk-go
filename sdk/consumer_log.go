
// consumer_log automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

type ConsumerLog struct {
    Id int `json:"id"`
    AppId int `json:"appId"`
    Ip string `json:"ip"`
    UserAgent string `json:"userAgent"`
    Method string `json:"method"`
    Path string `json:"path"`
    Header string `json:"header"`
    Body string `json:"body"`
    Date string `json:"date"`
}

