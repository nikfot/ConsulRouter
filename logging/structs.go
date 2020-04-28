package logging

import "time"

type reqDecoder interface{}

// Fielder interface is used for returning the fields to be logged, based on
// the API implementation. Different API calls can log different structured
// data, by implementing this interface.
type Fielder interface {
	Field() map[string]interface{}
}

// LogMessage is the basic structure used for log messages.
type LogMessage struct {
	Timestamp time.Time `json:"timestamp"`
	RequestID string    `json:"rid"` // Unique ID across end to end requests
	RTDelta   float64   `json:"rd"`  // Round Trip Delta, used to count the time until a request returns
	Endpoint  string    `json:"endpoint"`
	Request   string    `json:"request,omitempty"` // Return abstract structured data
	IP        string    `json:"ip,omitempty"`
	Message   string    `json:"message,omitempty"`
}
