// schema_update automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

type SchemaUpdate struct {
	Id       int          `json:"id"`
	Status   int          `json:"status"`
	Name     string       `json:"name"`
	Source   SchemaSource `json:"source"`
	Form     SchemaForm   `json:"form"`
	Metadata Metadata     `json:"metadata"`
}
