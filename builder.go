package gojsend

// JSendBuilder : interface for builder of JSend object
type JSendBuilder interface {
	Success(interface{}) JSendBuilder
	Fail(interface{}) JSendBuilder
	Error(string) JSendBuilder

	Data(interface{}) JSendBuilder
	Code(int) JSendBuilder
	Set(string, interface{}) JSendBuilder

	JSONEncoder(JSONEncoder) JSendBuilder
	Build() ([]byte, error)
	Response() map[string]interface{}
}

// JSendBuilderBuffer : JSendBuilder parameters buffer
type JSendBuilderBuffer struct {
	jsonEncoder JSONEncoder
	response    map[string]interface{}
}

// Set : sets custom key in the JSendBuilder
func (j *JSendBuilderBuffer) Set(key string, value interface{}) JSendBuilder {
	j.response[key] = value
	return j
}

// Success : sets status in the JSendBuilderBuffer to success
func (j *JSendBuilderBuffer) Success(data interface{}) JSendBuilder {
	j.Set(FieldStatus, StatusSuccess)
	j.Data(data)
	return j
}

// Fail : sets status in the JSendBuilderBuffer to fail
func (j *JSendBuilderBuffer) Fail(data interface{}) JSendBuilder {
	j.Set(FieldStatus, StatusFail)
	j.Data(data)
	return j
}

// Error : sets status in the JSendBuilderBuffer to error
func (j *JSendBuilderBuffer) Error(message string) JSendBuilder {
	j.Set(FieldStatus, StatusError)
	j.Set(FieldMessage, message)
	return j
}

// Data : sets data in the JSendBuilder
func (j *JSendBuilderBuffer) Data(data interface{}) JSendBuilder {
	return j.Set(FieldData, data)
}

// Code : sets code in the JSendBuilder
func (j *JSendBuilderBuffer) Code(code int) JSendBuilder {
	return j.Set(FieldCode, code)
}

// JSONEncoder : sets JSON encoder function to be used in the JSendBuilder
func (j *JSendBuilderBuffer) JSONEncoder(jsonEncoder JSONEncoder) JSendBuilder {
	j.jsonEncoder = jsonEncoder
	return j
}

// Response : returns the JSend response as map[string]interface{}
func (j *JSendBuilderBuffer) Response() map[string]interface{} {
	return j.response
}

// Build : encodes and returns JSend response as []byte
func (j *JSendBuilderBuffer) Build() ([]byte, error) {
	return j.jsonEncoder(j.response)
}
