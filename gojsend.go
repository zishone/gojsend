// Package gojsend is a JSend Builder and Response Writer for Go.
package gojsend

import (
	"encoding/json"
	"net/http"
)

// JSONEncoder : JSON encoder function
type JSONEncoder func(v interface{}) ([]byte, error)

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

// NewBuilder : returns JSendBuilder
func NewBuilder() JSendBuilder {
	return &JSendBuilderBuffer{
		jsonEncoder: json.Marshal,
		response:    make(map[string]interface{}),
	}
}

// NewWriter : returns JSendWriter which extends http.ResponseWriter with gojsend functions
func NewWriter(w http.ResponseWriter) JSendWriter {
	return &JSendWriterBuffer{
		builder:        NewBuilder(),
		statusCode:     http.StatusOK,
		responseWriter: w,
	}
}
