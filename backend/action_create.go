// action_create automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

type ActionCreate struct {
	Id       int          `json:"id"`
	Status   int          `json:"status"`
	Name     string       `json:"name"`
	Class    string       `json:"class"`
	Async    bool         `json:"async"`
	Engine   string       `json:"engine"`
	Config   ActionConfig `json:"config"`
	Metadata Metadata     `json:"metadata"`
}
