// event automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

type Event struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Schema      string   `json:"schema"`
	Metadata    Metadata `json:"metadata"`
}
