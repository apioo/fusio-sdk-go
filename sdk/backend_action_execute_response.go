
// backend_action_execute_response automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

// Represents an action execution response
type BackendActionExecuteResponse struct {
    StatusCode int `json:"statusCode"`
    Headers *BackendActionExecuteResponseHeaders `json:"headers"`
    Body *BackendActionExecuteResponseBody `json:"body"`
}

