package webserver

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	"nikfot/ConsulRouter/common"

	"nikfot/ConsulRouter/consulcalls"

	"nikfot/ConsulRouter/logging"
)

//ConsulRouter receives calls to be routed to consul APIS
func ConsulRouter(w http.ResponseWriter, r *http.Request) {

	// Record the time to calculate the delta
	tStart := time.Now()
	// Get unique ID for the request (used for logging)
	uuid := logging.NewUUID()

	defer r.Body.Close()
	//if request was made to the UI ("/ui") or default page("/") serve static page
	if r.URL.Path[1:] == "" {
		uiHandler(w, r)
	} else {
		resp := consulcalls.RouteToConsul(r)

		fmt.Printf("API call, %s, routed to Consul API \n", r.URL.Path[1:])
		fmt.Fprintf(w, resp)
		ip := common.GetIP(r)
		logging.Log(uuid, tStart, r.RequestURI, r.Method, ip, common.ResultSuccess)
		logging.LogInfo("INFO", common.TruncateString(string(resp), 500))
	}

}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	// Record the time to calculate the delta
	tStart := time.Now()
	// Get unique ID for the request (used for logging)
	uuid := logging.NewUUID()
	ip := common.GetIP(r)
	http.ServeFile(w, r, "misc/favicon.ico")
	logging.Log(uuid, tStart, r.RequestURI, r.Method, ip, "FAVICON")

}

func uiHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("misc/index.html"))
	data := ServiceDiscoveryUI{
		PageTitle: "Service Discovery",
		Info:      "This is unified service discovery powered by Consul.",
	}
	tmpl.Execute(w, data)

}
