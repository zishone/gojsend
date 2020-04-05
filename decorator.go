package gojsend

import "net/http"

// JSendDecorator : interface for decorator of http.HandlerFunc
type JSendDecorator interface {
	JSONEncoder(JSONEncoder) JSendDecorator
	Decorate(JSendHandler) http.HandlerFunc
}

// JSendDecoratorBuffer : JSendDecorator parameters buffer
type JSendDecoratorBuffer struct {
	jsonEncoder JSONEncoder
}

// JSONEncoder : sets JSON encoder function to be used in the
// JSendDecoratorBuffer
func (j *JSendDecoratorBuffer) JSONEncoder(jsonEncoder JSONEncoder) JSendDecorator {
	j.jsonEncoder = jsonEncoder
	return j
}

// Decorate : decorates the handler to replace http.ResponseWrtier with
// JSendWriter
func (j *JSendDecoratorBuffer) Decorate(h JSendHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jw := NewWriter(w).
			JSONEncoder(j.jsonEncoder)
		h(jw, r)
	}
}
