
// backend_dashboard_transaction automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import "time"
type BackendDashboardTransaction struct {
    Id int `json:"id"`
    Status string `json:"status"`
    Provider string `json:"provider"`
    TransactionId string `json:"transactionId"`
    Amount float64 `json:"amount"`
    Date time.Time `json:"date"`
}
