// backend_audit automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package fusio

import "time"

type BackendAudit struct {
	Id      int                `json:"id"`
	App     BackendApp         `json:"app"`
	User    BackendUser        `json:"user"`
	Event   string             `json:"event"`
	Ip      string             `json:"ip"`
	Message string             `json:"message"`
	Content BackendAuditObject `json:"content"`
	Date    time.Time          `json:"date"`
}
