package gojsend

import (
	"net/http"
)

// JSendWriter : extends http.ResponseWriter to add gojsend functions
type JSendWriter interface {
	http.ResponseWriter

	Success(interface{}) JSendWriter
	Fail(interface{}) JSendWriter
	Error(string) JSendWriter

	Data(interface{}) JSendWriter
	Code(int) JSendWriter
	Set(string, interface{}) JSendWriter

	JSONEncoder(JSONEncoder) JSendWriter

	Send() (int, error)
}

// JSendExtenderBuffer : JSendWriter parameters buffer
type JSendExtenderBuffer struct {
	builder        JSendBuilder
	statusCode     int
	responseWriter http.ResponseWriter
}

// ExtendWriter : returns JSendWriter which extends http.ResponseWriter with gojsend functions
func ExtendWriter(w http.ResponseWriter) JSendWriter {
	return &JSendExtenderBuffer{
		builder:        NewBuilder(),
		statusCode:     StatusCodeSuccess,
		responseWriter: w,
	}
}

// Set : sets custom key in the JSendExtenderBuffer
func (j *JSendExtenderBuffer) Set(key string, value interface{}) JSendWriter {
	j.builder.Set(key, value)
	return j
}

// Success : sets status in the JSendExtenderBuffer to success
func (j *JSendExtenderBuffer) Success(data interface{}) JSendWriter {
	j.statusCode = StatusCodeSuccess
	j.builder.Success(data)
	return j
}

// Fail : sets status in the JSendExtenderBuffer to fail
func (j *JSendExtenderBuffer) Fail(data interface{}) JSendWriter {
	j.statusCode = StatusCodeFail
	j.builder.Fail(data)
	return j
}

// Error : sets status in the JSendExtenderBuffer to error
func (j *JSendExtenderBuffer) Error(message string) JSendWriter {
	j.statusCode = StatusCodeError
	j.builder.Error(message)
	return j
}

// Data : sets data in the JSendExtenderBuffer
func (j *JSendExtenderBuffer) Data(data interface{}) JSendWriter {
	j.builder.Data(data)
	return j
}

// Code : sets code in the JSendExtenderBuffer
func (j *JSendExtenderBuffer) Code(code int) JSendWriter {
	j.builder.Code(code)
	return j
}

// JSONEncoder : sets JSON encoder function to be used in the JSendExtenderBuffer
func (j *JSendExtenderBuffer) JSONEncoder(jsonEncoder JSONEncoder) JSendWriter {
	j.builder.JSONEncoder(jsonEncoder)
	return j
}

// Send : encodes and sends response
func (j *JSendExtenderBuffer) Send() (int, error) {
	bs, err := j.builder.Build()
	if err != nil {
		return 0, err
	}
	if j.Header().Get("Content-Type") == "" {
		j.Header().Set("Content-Type", "application/json")
	}
	j.WriteHeader(j.statusCode)
	return j.Write(bs)
}
