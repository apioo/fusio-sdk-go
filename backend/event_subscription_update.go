// event_subscription_update automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

type EventSubscriptionUpdate struct {
	Id        int                         `json:"id"`
	EventId   int                         `json:"eventId"`
	UserId    int                         `json:"userId"`
	Endpoint  string                      `json:"endpoint"`
	Responses []EventSubscriptionResponse `json:"responses"`
}
