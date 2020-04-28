package webserver

import "net/http"

var requestRouter *http.ServeMux

//ServiceDiscoveryUI handles UI elements
type ServiceDiscoveryUI struct {
	PageTitle string
	Info      string
}
