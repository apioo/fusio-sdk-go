// common_form_element_select automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package sdk

type CommonFormElementSelect struct {
	Element string                          `json:"element"`
	Name    string                          `json:"name"`
	Title   string                          `json:"title"`
	Help    string                          `json:"help"`
	Options []CommonFormElementSelectOption `json:"options"`
}
