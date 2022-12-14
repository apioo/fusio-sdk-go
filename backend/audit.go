// audit automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

import "time"

type Audit struct {
	Id      int         `json:"id"`
	App     App         `json:"app"`
	User    User        `json:"user"`
	Event   string      `json:"event"`
	Ip      string      `json:"ip"`
	Message string      `json:"message"`
	Content AuditObject `json:"content"`
	Date    time.Time   `json:"date"`
}
