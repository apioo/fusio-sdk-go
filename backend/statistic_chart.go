// statistic_chart automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

type StatisticChart struct {
	Labels []string             `json:"labels"`
	Data   []StatisticChartData `json:"data"`
	Series []string             `json:"series"`
}
