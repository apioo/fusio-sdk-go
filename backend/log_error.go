// log_error automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app

package backend

type LogError struct {
	Id      int    `json:"id"`
	Message int    `json:"message"`
	Trace   string `json:"trace"`
	File    string `json:"file"`
	Line    int    `json:"line"`
}
