
// backend_rate_update automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import "time"
type BackendRateUpdate struct {
    Id int `json:"id"`
    Priority int `json:"priority"`
    Name string `json:"name"`
    RateLimit int `json:"rateLimit"`
    Timespan time.Duration `json:"timespan"`
    Allocation []BackendRateAllocation `json:"allocation"`
    Metadata *CommonMetadata `json:"metadata"`
}
