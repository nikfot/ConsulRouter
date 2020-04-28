package configuration

import (
	"os"
)

//LoadConfiguration loads the configuration from env variables
func LoadConfiguration() {
	consulServers, port, host := readConfiguration()

	var serverName string
	var serial int
	for i := range consulServers {
		serial = i + 1
		serverName = "CONSUL-SERVER-" + string(serial)
		os.Setenv(serverName, consulServers[i])
	}
	os.Setenv("PORT", port)
	os.Setenv("HOST", host)

}
