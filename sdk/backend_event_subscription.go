
// backend_event_subscription automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk
type BackendEventSubscription struct {
    Id int `json:"id"`
    EventId int `json:"eventId"`
    UserId int `json:"userId"`
    Endpoint string `json:"endpoint"`
    Responses []BackendEventSubscriptionResponse `json:"responses"`
}
