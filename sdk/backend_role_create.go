
// backend_role_create automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

type BackendRoleCreate struct {
    Id int `json:"id"`
    CategoryId int `json:"categoryId"`
    Name string `json:"name"`
    Scopes []string `json:"scopes"`
}

