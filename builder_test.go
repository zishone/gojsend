package gojsend_test

import (
	"encoding/json"
	"testing"

	"github.com/zishone/gojsend"
)

func TestNewBuilder(t *testing.T) {
	j := gojsend.NewBuilder()

	j, ok := j.(gojsend.JSendBuilder)
	if !ok {
		t.Errorf("NewBuilder\n\thave: %T\n\twant: %v", j, "JSendBuilder")
	}
}

func TestBuilderSet(t *testing.T) {
	key := "foo"
	value := "bar"
	j, err := gojsend.NewBuilder().Set(key, value).Build()
	if err != nil {
		t.Errorf("BuilderSet\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(j, &body)
	if err != nil {
		t.Errorf("BuilderSet\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[key] != value {
		t.Errorf("BuilderSet\n\thave: %v\n\twant: %v", body[key], value)
	}
}

func TestBuilderSuccess(t *testing.T) {
	key := "foo"
	value := "bar"
	j, err := gojsend.NewBuilder().Success(map[string]interface{}{key: value}).Build()
	if err != nil {
		t.Errorf("BuilderSuccess\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(j, &body)
	if err != nil {
		t.Errorf("BuilderSuccess\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[gojsend.FieldStatus] != gojsend.StatusSuccess {
		t.Errorf("BuilderSuccess\n\thave: %v\n\twant: %v", body[gojsend.FieldStatus], gojsend.StatusSuccess)
	}
	data, ok := body[gojsend.FieldData].(map[string]interface{})
	if !ok {
		t.Errorf("BuilderSuccess\n\thave: %T\n\twant: %v", data, "map[string]interface{}")
	}
	if data[key] != value {
		t.Errorf("BuilderSuccess\n\thave: %v\n\twant: %v", data[key], value)
	}
}

func TestBuilderFail(t *testing.T) {
	key := "foo"
	value := "bar"
	j, err := gojsend.NewBuilder().Fail(map[string]interface{}{key: value}).Build()
	if err != nil {
		t.Errorf("BuilderFail\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(j, &body)
	if err != nil {
		t.Errorf("BuilderFail\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[gojsend.FieldStatus] != gojsend.StatusFail {
		t.Errorf("BuilderFail\n\thave: %v\n\twant: %v", body[gojsend.FieldStatus], gojsend.StatusFail)
	}
	data, ok := body[gojsend.FieldData].(map[string]interface{})
	if !ok {
		t.Errorf("BuilderFail\n\thave: %T\n\twant: %v", data, "map[string]interface{}")
	}
	if data[key] != value {
		t.Errorf("BuilderFail\n\thave: %v\n\twant: %v", data[key], value)
	}
}

func TestBuilderError(t *testing.T) {
	message := "message"
	j, err := gojsend.NewBuilder().Error(message).Build()
	if err != nil {
		t.Errorf("BuilderError\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(j, &body)
	if err != nil {
		t.Errorf("BuilderError\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[gojsend.FieldStatus] != gojsend.StatusError {
		t.Errorf("BuilderError\n\thave: %v\n\twant: %v", body[gojsend.FieldStatus], gojsend.StatusError)
	}
	if body[gojsend.FieldMessage] != message {
		t.Errorf("BuilderError\n\thave: %v\n\twant: %v", body[gojsend.FieldMessage], message)
	}
}

func TestBuilderData(t *testing.T) {
	key := "foo"
	value := "bar"
	j, err := gojsend.NewBuilder().Data(map[string]interface{}{key: value}).Build()
	if err != nil {
		t.Errorf("BuilderData\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(j, &body)
	if err != nil {
		t.Errorf("BuilderData\n\thave: %v\n\twant: %v", err, nil)
	}
	data, ok := body[gojsend.FieldData].(map[string]interface{})
	if !ok {
		t.Errorf("BuilderData\n\thave: %T\n\twant: %v", data, "map[string]interface{}")
	}
	if data[key] != value {
		t.Errorf("BuilderData\n\thave: %v\n\twant: %v", data[key], value)
	}
}

func TestBuilderCode(t *testing.T) {
	code := 0
	j, err := gojsend.NewBuilder().Code(code).Build()
	if err != nil {
		t.Errorf("BuilderCode\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(j, &body)
	if err != nil {
		t.Errorf("BuilderCode\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[gojsend.FieldCode] != float64(code) {
		t.Errorf("BuilderCode\n\thave: %v\n\twant: %v", body[gojsend.FieldCode], code)
	}
}

func TestBuilderJSONEncoder(t *testing.T) {
	j, err := gojsend.NewBuilder().JSONEncoder(json.Marshal).Build()
	if err != nil {
		t.Errorf("BuilderJSONEncoder\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(j, &body)
	if err != nil {
		t.Errorf("BuilderJSONEncoder\n\thave: %v\n\twant: %v", err, nil)
	}
}

func TestBuilderBuild(t *testing.T) {
	j, err := gojsend.NewBuilder().Build()
	if err != nil {
		t.Errorf("BuilderBuild\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.Unmarshal(j, &body)
	if err != nil {
		t.Errorf("BuilderBuild\n\thave: %v\n\twant: %v", err, nil)
	}
}
