// system_about automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package fusio

type SystemAbout struct {
	ApiVersion      string            `json:"apiVersion"`
	Title           string            `json:"title"`
	Description     string            `json:"description"`
	TermsOfService  string            `json:"termsOfService"`
	ContactName     string            `json:"contactName"`
	ContactUrl      string            `json:"contactUrl"`
	ContactEmail    string            `json:"contactEmail"`
	LicenseName     string            `json:"licenseName"`
	LicenseUrl      string            `json:"licenseUrl"`
	PaymentCurrency string            `json:"paymentCurrency"`
	Categories      []string          `json:"categories"`
	Scopes          []string          `json:"scopes"`
	Apps            SystemAboutApps   `json:"apps"`
	Links           []SystemAboutLink `json:"links"`
}
