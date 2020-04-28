package common

const (
	// ResultSuccess defines the string returned in cases of success.
	ResultSuccess = "SUCCESS"

	// ResultCreated defines the string returned in cases of created.
	ResultCreated = "SUCCESS"

	// ResultError defines the string returned in cases of error.
	ResultError = "TECHNICAL_ERROR"

	// ResultUnauthorized defines the string returned in cases of error.
	ResultUnauthorized = "UNAUTHORIZED"

	// ResultFailed defines the string returned in cases of failed requests.
	ResultFailed = "FAILED"

	// ResultMethodNotAllowed defines the string returned in cases of wromg method.
	ResultMethodNotAllowed = "METHOD_NOT_ALLOWED"

	// DefaultTimeout defines the default timeout used in Context, when
	// a timeout has not been specified in the request.
	DefaultTimeout = "10s"
)
