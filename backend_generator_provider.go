// backend_generator_provider automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package fusio

type BackendGeneratorProvider struct {
	Path   string                         `json:"path"`
	Scopes []string                       `json:"scopes"`
	Public bool                           `json:"public"`
	Config BackendGeneratorProviderConfig `json:"config"`
}
