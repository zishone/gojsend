// Package gojsend provides interfaces for building and writing of JSON
// responses as defined by JSend (https://github.com/omniti-labs/jsend).
//
//
// JSendBuilder
// The JSend Builder provides a convinient way for building JSend
// responses.
// Examples:
//
//  // Success
//  builder := gojsend.NewBuilder().
//  	Success(map[string]interface{}{"foo": "bar"})
//  b := string(builder.Build()) 	// {"status":"success","data":{"foo":"bar"}}
//  m := builder.Response()				// map[status:success data:map[foo:bar]]
//
//  // Fail
//  builder := gojsend.NewBuilder().
//  	Fail(map[string]interface{}{"foo": "bar"})
//  b := string(builder.Build()) 	// {"status":"fail","data":{"foo":"bar"}}
//  m := builder.Response() 			// map[status:fail data:map[foo:bar]]
//
//  // Error
//  builder := gojsend.NewBuilder().
//  	Error("foobar").
//  	Code(1).
//  	Data(map[string]interface{}{"foo": "bar"})
//  b := string(builder.Build()) 	// {"status":"error","message":"foobar","code":1,"data":{"foo":"bar"}}
//  m := builder.Response() 			// map[status:error message:foobar code:1 data:map[foo:bar]]
//
//
// JSendWriter
// The JSend Writer extends the http.ResponseWriter with the JSend
// builder functions and a Send function to pass the built JSend
// response to http.ResponseWriter's Write function.
// Examples:
//
//  // Success
//	func HandlerFunc(w http.ResponseWriter, r *http.Request) {
//  	gojsend.NewWriter(w).
//  		Success(map[string]interface{}{"foo": "bar"}).
//			// StatusCode(200). - Use to overwrite default status code
//			Send() 	// Response Body is the same with Builder example with HTTP Status Code 200
//	}
//
//  // Fail
//	func HandlerFunc(w http.ResponseWriter, r *http.Request) {
//  	gojsend.NewWriter(w).
//  		Fail(map[string]interface{}{"foo": "bar"}).
//      // StatusCode(200). - Use to overwrite default status code
//			Send() 	// Response Body is the same with Builder example with HTTP Status Code 400
//	}
//
//  // Error
//	func HandlerFunc(w http.ResponseWriter, r *http.Request) {
//  	gojsend.NewWriter(w).
//  		Error("foobar").
//  		Code(1).
//  		Data(map[string]interface{}{"foo": "bar"}).
//      // StatusCode(200). - Use to overwrite default status code
//			Send() 	// Response Body is the same with Builder example with HTTP Status Code 500
//	}
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
