// dashboard_request automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

import "time"

type DashboardRequest struct {
	Id   int       `json:"id"`
	Path string    `json:"path"`
	Ip   string    `json:"ip"`
	Date time.Time `json:"date"`
}
