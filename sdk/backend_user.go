
// backend_user automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

type BackendUser struct {
    Id int `json:"id"`
    RoleId int `json:"roleId"`
    PlanId int `json:"planId"`
    Status int `json:"status"`
    Name string `json:"name"`
    Email string `json:"email"`
    Points int `json:"points"`
    Scopes []string `json:"scopes"`
    Apps []BackendApp `json:"apps"`
    Metadata *CommonMetadata `json:"metadata"`
    Date string `json:"date"`
}

