package gojsend

import "net/http"

// Header : implements http.ResponseWriter's Header() Header function
func (j *JSendWriterBuffer) Header() http.Header {
	return j.responseWriter.Header()
}

// Header : implements http.ResponseWriter's Write([]byte) (int, error) function
func (j *JSendWriterBuffer) Write(b []byte) (int, error) {
	return j.responseWriter.Write(b)
}

// WriteHeader : implements http.ResponseWriter's WriteHeader(statusCode int) function
func (j *JSendWriterBuffer) WriteHeader(statusCode int) {
	j.responseWriter.WriteHeader(statusCode)
}
