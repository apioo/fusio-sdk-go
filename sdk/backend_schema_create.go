// backend_schema_create automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

type BackendSchemaCreate struct {
	Id       int                 `json:"id"`
	Status   int                 `json:"status"`
	Name     string              `json:"name"`
	Source   BackendSchemaSource `json:"source"`
	Form     BackendSchemaForm   `json:"form"`
	Metadata CommonMetadata      `json:"metadata"`
}
