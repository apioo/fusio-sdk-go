// plan_usage_collection_query automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

import "time"

type PlanUsageCollectionQuery struct {
	StartIndex int       `json:"startIndex"`
	Count      int       `json:"count"`
	Search     string    `json:"search"`
	From       time.Time `json:"from"`
	To         time.Time `json:"to"`
	RouteId    int       `json:"routeId"`
	AppId      int       `json:"appId"`
	UserId     int       `json:"userId"`
}
