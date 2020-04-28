package health

var (
	// Version -- the semantic version of the app
	Version string
	// Build -- the last git commit
	Build string
	// Date -- the date of the commit
	Date string
)

// statusMapResponse is a structure that is used as part in every map response.
type statusMapResponse struct {
	MapResponse
}

// MapResponse is a structure that is used as part in every map response.
type MapResponse struct {
	Result      string            `json:"result"`
	Description map[string]string `json:"description,omitempty"`
}
