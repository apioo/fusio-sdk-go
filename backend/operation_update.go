// operation_update automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

type OperationUpdate struct {
	Id          int                 `json:"id"`
	Status      int                 `json:"status"`
	Active      bool                `json:"active"`
	Public      bool                `json:"public"`
	Stability   int                 `json:"stability"`
	Description string              `json:"description"`
	HttpMethod  string              `json:"httpMethod"`
	HttpPath    string              `json:"httpPath"`
	HttpCode    int                 `json:"httpCode"`
	Name        string              `json:"name"`
	Parameters  OperationParameters `json:"parameters"`
	Incoming    string              `json:"incoming"`
	Outgoing    string              `json:"outgoing"`
	Throws      OperationThrows     `json:"throws"`
	Action      string              `json:"action"`
	Costs       int                 `json:"costs"`
	Scopes      []string            `json:"scopes"`
	Metadata    Metadata            `json:"metadata"`
}