package constants

type HeaderKeys string

const (
	XCorrelationID HeaderKeys = "x-correlation-id"
	XRequestID     HeaderKeys = "x-request-id"
	XTraceID       HeaderKeys = "x-trace-id"
	XSpanID        HeaderKeys = "x-span-id"
	XAppVersion    HeaderKeys = "x-app-version"
)
