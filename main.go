package main

<<<<<<< HEAD
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
=======
import "net/http"

func main() {
	http.HandleFunc("/", webserver.ConsulRouter)
	http.ListenAndServe(":8080", nil)

>>>>>>> Initial Commit
}
