/*
Package logging provides the login mechanisms for the application/server.
*/
package logging

import (
	"fmt"
	"log"
	"strings"
	"time"

	"nikfot/ConsulRouter/common"

	uuid "github.com/satori/go.uuid"
)

//LogSetup Initializes logging
func LogSetup() {

	envVars := common.GetEnvVars()

	log.Printf("Server will run on host: %s , port: %s\n", envVars["HOST"], envVars["PORT"])
	serverName := "CONSUL-SERVER-"
	serial := 1
	key := serverName + string(serial)
	for range envVars {
		if common.IsEnvExist(key) {
			log.Printf("Redirecting to Consul url: %s\n", envVars[key])
		}
		serial += serial
		key = serverName + string(serial)
	}

}

//LogRequestPayload logs information on proxy forwarding
func LogRequestPayload(consulURL string) {
	log.Printf("call: v ,proxy_url: %s\n", consulURL)
}

// Log is used to log a json structured message with LogMessage data
func Log(id string, tStart time.Time, endpoint string, req string, ip string, message string) {
	lg := &LogMessage{}
	lg.Timestamp = time.Now()
	if id != "" {
		lg.RequestID = id
	}
	lg.RTDelta = time.Since(tStart).Seconds()
	lg.Endpoint = endpoint
	lg.Message = message
	lg.IP = ip
	if req != "" {
		lg.Request = req
	}

	fmt.Printf("%s: INFO - %s %f %s %+v %s %s\n", lg.Timestamp, lg.RequestID, lg.RTDelta, lg.Endpoint, lg.Request, lg.IP, lg.Message)
}

// LogInfo returns just message for debugging
func LogInfo(level string, message string) {
	fmt.Printf("%s: %s - %s\n", time.Now(), strings.ToUpper(level), message)
}

// NewUUID functions generates and returns a UUIDv4 string
func NewUUID() string {
	u, _ := uuid.NewV4()
	return u.String()
	//return uuid.NewV4().String()
}
