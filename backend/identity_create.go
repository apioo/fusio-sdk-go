// identity_create automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

type IdentityCreate struct {
	Id          int            `json:"id"`
	AppId       int            `json:"appId"`
	RoleId      int            `json:"roleId"`
	Name        string         `json:"name"`
	Icon        string         `json:"icon"`
	Class       string         `json:"class"`
	Config      IdentityConfig `json:"config"`
	AllowCreate bool           `json:"allowCreate"`
}