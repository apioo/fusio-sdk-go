// statistic_count automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

import "time"

type StatisticCount struct {
	Count int       `json:"count"`
	From  time.Time `json:"from"`
	To    time.Time `json:"to"`
}
