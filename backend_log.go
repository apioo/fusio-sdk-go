// backend_log automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package fusio

import "time"

type BackendLog struct {
	Id        int               `json:"id"`
	Ip        string            `json:"ip"`
	UserAgent string            `json:"userAgent"`
	Method    string            `json:"method"`
	Path      string            `json:"path"`
	Header    string            `json:"header"`
	Body      string            `json:"body"`
	Date      time.Time         `json:"date"`
	Errors    []BackendLogError `json:`
}
