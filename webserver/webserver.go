package webserver

import (
	"log"
	"net/http"
	"os"

	"nikfot/ConsulRouter/health"
)

//Start initiates the webserver
func Start() {
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	log.Println("Starting HTTP server(host: " + host + ", port: " + port + ")")
	if err := http.ListenAndServe(host+":"+port, nil); err != nil {
		panic(err)
	}
}

//LoadEndpoints loads the handlers of each endpoint
func LoadEndpoints() {
	if requestRouter == nil {
		requestRouter = http.NewServeMux()
	}

	http.HandleFunc("/", ConsulRouter)
	http.HandleFunc("/ui/", uiHandler)
	http.HandleFunc("/liveness", health.LivenessGet)
	http.HandleFunc("/readiness", health.ReadinessGet)
	http.HandleFunc("/version", health.GetVersion)
	http.HandleFunc("/favicon.ico", faviconHandler)

}
