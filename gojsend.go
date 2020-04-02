package gojsend

// JSONEncoder : JSON encoder function
type JSONEncoder func(v interface{}) ([]byte, error)

const (
	// StatusCodeSuccess : HTTP success status code
	StatusCodeSuccess = 200
	// StatusCodeFail : HTTP fail status code
	StatusCodeFail = 400
	// StatusCodeError : HTTP error status code
	StatusCodeError = 500
)

const (
	// StatusSuccess : JSend success status
	StatusSuccess = "success"
	// StatusFail : JSend fail status
	StatusFail = "fail"
	// StatusError : JSend error status
	StatusError = "error"
)

const (
	// FieldStatus : JSend status field
	FieldStatus = "status"
	// FieldData : JSend data field
	FieldData = "data"
	// FieldMessage : JSend message field
	FieldMessage = "message"
	// FieldCode : JSend code field
	FieldCode = "code"
)
