package main

import (
	"nikfot/ConsulRouter/logging"

	"nikfot/ConsulRouter/webserver"

	"nikfot/ConsulRouter/configuration"
)

func main() {
	configuration.LoadConfiguration()
	logging.LogSetup()
	webserver.LoadEndpoints()
	webserver.Start()
}
