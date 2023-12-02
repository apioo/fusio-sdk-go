// backend_event_subscription_response automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

import "time"

type BackendEventSubscriptionResponse struct {
	Id          int       `json:"id"`
	Status      int       `json:"status"`
	Code        int       `json:"code"`
	Attempts    int       `json:"attempts"`
	Error       string    `json:"error"`
	ExecuteDate time.Time `json:"executeDate"`
}
