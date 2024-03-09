
// consumer_transaction automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

import "time"
type ConsumerTransaction struct {
    Id int `json:"id"`
    UserId int `json:"userId"`
    PlanId int `json:"planId"`
    Plan ConsumerPlan `json:"plan"`
    TransactionId string `json:"transactionId"`
    Amount float64 `json:"amount"`
    Points float64 `json:"points"`
    PeriodStart time.Time `json:"periodStart"`
    PeriodEnd time.Time `json:"periodEnd"`
    InsertDate time.Time `json:"insertDate"`
}
