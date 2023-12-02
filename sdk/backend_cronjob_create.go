// backend_cronjob_create automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

import "time"

type BackendCronjobCreate struct {
	Id          int                   `json:"id"`
	Name        string                `json:"name"`
	Cron        string                `json:"cron"`
	Action      string                `json:"action"`
	ExecuteDate time.Time             `json:"executeDate"`
	ExitCode    int                   `json:"exitCode"`
	Metadata    CommonMetadata        `json:"metadata"`
	Errors      []BackendCronjobError `json:`
}
