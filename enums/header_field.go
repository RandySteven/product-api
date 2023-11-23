package enums

type ContextHeaderKey string

const (
	ContentType ContextHeaderKey = "Content-Type"
	XRequestID                   = "x-request-id"
	XTimestamp                   = "x-timestamp"
)
