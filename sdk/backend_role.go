// backend_role automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

type BackendRole struct {
	Id         int      `json:"id"`
	CategoryId int      `json:"categoryId"`
	Name       string   `json:"name"`
	Scopes     []string `json:"scopes"`
}
