
// consumer_webhook automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk
type ConsumerWebhook struct {
    Id int `json:"id"`
    Status int `json:"status"`
    Event string `json:"event"`
    Name string `json:"name"`
    Endpoint string `json:"endpoint"`
    Responses []ConsumerWebhookResponse `json:"responses"`
}
