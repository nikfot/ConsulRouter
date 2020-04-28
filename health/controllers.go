package health

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"nikfot/ConsulRouter/common"
	"nikfot/ConsulRouter/consulcalls"
	"nikfot/ConsulRouter/logging"
)

//LivenessGet is used to check if app is live
func LivenessGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

//ReadinessGet check if connection is ok
func ReadinessGet(w http.ResponseWriter, r *http.Request) {
	if consulcalls.IsHealthy() {
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Readiness probe is not OK", http.StatusServiceUnavailable)
	}
}

// GetVersion returns the version description of the application
func GetVersion(w http.ResponseWriter, r *http.Request) {
	// Record the time to calculate the delta
	tStart := time.Now()
	// Get unique ID for the request (used for logging)
	uuid := logging.NewUUID()
	versionOut := make(map[string]string)
	versionOut["Version"] = Version
	versionOut["Revision"] = Build
	versionOut["CommitTime"] = Date
	response := statusMapResponse{
		MapResponse: MapResponse{
			Result:      common.ResultSuccess,
			Description: versionOut},
	}
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusSuccess)
	resp, _ := json.Marshal(response)
	//json.NewEncoder(w).Encode("wever")
	fmt.Fprintf(w, string(resp))
	//json.NewEncoder(w).Encode("develop")
	ip := common.GetIP(r)
	logging.Log(uuid, tStart, r.RequestURI, r.Method, ip, common.ResultSuccess)
}
