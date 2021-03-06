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
	Builder() JSendBuilder

	StatusCode(int) JSendWriter
	Send() (int, error)
}

// JSendWriterBuffer : JSendWriter parameters buffer
type JSendWriterBuffer struct {
	builder        JSendBuilder
	statusCode     int
	responseWriter http.ResponseWriter
}

// Set : sets custom key in the JSendWriterBuffer
func (j *JSendWriterBuffer) Set(key string, value interface{}) JSendWriter {
	j.builder.Set(key, value)
	return j
}

// Success : sets status in the JSendWriterBuffer to success
func (j *JSendWriterBuffer) Success(data interface{}) JSendWriter {
	j.StatusCode(http.StatusOK)
	j.builder.Success(data)
	return j
}

// Fail : sets status in the JSendWriterBuffer to fail
func (j *JSendWriterBuffer) Fail(data interface{}) JSendWriter {
	j.StatusCode(http.StatusBadRequest)
	j.builder.Fail(data)
	return j
}

// Error : sets status in the JSendWriterBuffer to error
func (j *JSendWriterBuffer) Error(message string) JSendWriter {
	j.StatusCode(http.StatusInternalServerError)
	j.builder.Error(message)
	return j
}

// Data : sets data in the JSendWriterBuffer
func (j *JSendWriterBuffer) Data(data interface{}) JSendWriter {
	j.builder.Data(data)
	return j
}

// Code : sets code in the JSendWriterBuffer
func (j *JSendWriterBuffer) Code(code int) JSendWriter {
	j.builder.Code(code)
	return j
}

// JSONEncoder : sets JSON encoder function to be used in the
// JSendWriterBuffer
func (j *JSendWriterBuffer) JSONEncoder(jsonEncoder JSONEncoder) JSendWriter {
	j.builder.JSONEncoder(jsonEncoder)
	return j
}

// StatusCode : sets HTTP status code in the JSendWriterBuffer
func (j *JSendWriterBuffer) StatusCode(statusCode int) JSendWriter {
	j.statusCode = statusCode
	return j
}

// Builder : returns the JSend builder in the JSendWriterBuffer
func (j *JSendWriterBuffer) Builder() JSendBuilder {
	return j.builder
}

// Send : encodes and sends response
func (j *JSendWriterBuffer) Send() (int, error) {
	bs, err := j.Builder().Build()
	if err != nil {
		return 0, err
	}
	if j.Header().Get("Content-Type") == "" {
		j.Header().Set("Content-Type", "application/json")
	}
	j.WriteHeader(j.statusCode)
	return j.Write(bs)
}
