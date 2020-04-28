package consulcalls

import "net/http"

type reqDecoder interface{}

//Registry is a map o services, key is the service name and value the targets of the service
type Registry map[string][]string

//ConsulRegistry is a registry of Consul APIs
var ConsulRegistry = Registry{
	"consulHost": {
		"http://lh-consul.lab.up",
		"http://mn-consul.lab.up",
	},
}

// Response is a structure that is used as part in every response.
// All responses should include this struct.
type Response struct {
	Result      string `json:"result"`
	Description string `json:"description,omitempty"`
}

type transport struct {
	http.RoundTripper
}

// ForwardRequest represents a valid HTTP response
type ForwardRequest struct {
	Request *http.Request
}

// ForwardResponse represents a valid HTTP response
type ForwardResponse struct {
	Response
}
