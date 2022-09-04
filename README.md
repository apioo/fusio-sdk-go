
# Fusio SDK

This Go library helps to talk to the Fusio (https://www.fusio-project.org/)
API. It uses the automatically generated definition of the Fusio backend.

## Usage

The following example shows how you can get a client for the backend to create
a new action at the backend.

```go
func main() {
	var store = &sdkgen.MemoryTokenStore{}
	var scopes = []string{"backend"}

	var client = NewClient("", "", "", store, scopes)

	var config = backend.ActionConfig{}
	config["event"] = "my_custom_event"

	var data = backend.ActionCreate{
		Status: 1,
		Async:  false,
		Name:   "My_New_action",
		Class:  "Fusio\\Adapter\\Util\\Action\\UtilDispatchEvent",
		Config: config,
	}

	message, error := client.Backend.BackendAction().GetBackendAction().BackendActionActionCreate(data)

	if error != nil {
		log.Panic(error)
	}

	if message.Success == false {
		log.Panic("Could not create action server responded with an error: " + message.Message)
	}

	if message.Success == true {
		print("Created a new action at the backend")
	}
}
```
