package gojsend_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/zishone/gojsend"
)

func TestExtendWriter(t *testing.T) {
	w := httptest.NewRecorder()
	j := gojsend.ExtendWriter(w)
	j, ok := j.(gojsend.JSendWriter)
	if !ok {
		t.Errorf("ExtendWriter\n\thave: %T\n\twant: %v", j, "JSendWriter")
	}
}

func TestExtenderSet(t *testing.T) {
	key := "foo"
	value := "bar"
	w := httptest.NewRecorder()
	if _, err := gojsend.ExtendWriter(w).Set(key, value).Send(); err != nil {
		t.Errorf("ExtenderSet\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err := json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("ExtenderSet\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[key] != value {
		t.Errorf("ExtenderSet\n\thave: %v\n\twant: %v", body[key], value)
	}
}

func TestExtenderSuccess(t *testing.T) {
	key := "foo"
	value := "bar"
	w := httptest.NewRecorder()
	_, err := gojsend.ExtendWriter(w).Success(map[string]interface{}{key: value}).Send()
	if err != nil {
		t.Errorf("ExtenderSuccess\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("ExtenderSuccess\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[gojsend.FieldStatus] != gojsend.StatusSuccess {
		t.Errorf("ExtenderSuccess\n\thave: %v\n\twant: %v", body[gojsend.FieldStatus], gojsend.StatusSuccess)
	}
	data, ok := body[gojsend.FieldData].(map[string]interface{})
	if !ok {
		t.Errorf("ExtenderSuccess\n\thave: %T\n\twant: %v", data, "map[string]interface{}")
	}
	if data[key] != value {
		t.Errorf("ExtenderSuccess\n\thave: %v\n\twant: %v", data[key], value)
	}
}

func TestExtenderFail(t *testing.T) {
	key := "foo"
	value := "bar"
	w := httptest.NewRecorder()
	_, err := gojsend.ExtendWriter(w).Fail(map[string]interface{}{key: value}).Send()
	if err != nil {
		t.Errorf("ExtenderFail\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("ExtenderFail\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[gojsend.FieldStatus] != gojsend.StatusFail {
		t.Errorf("ExtenderFail\n\thave: %v\n\twant: %v", body[gojsend.FieldStatus], gojsend.StatusFail)
	}
	data, ok := body[gojsend.FieldData].(map[string]interface{})
	if !ok {
		t.Errorf("ExtenderFail\n\thave: %T\n\twant: %v", data, "map[string]interface{}")
	}
	if data[key] != value {
		t.Errorf("ExtenderFail\n\thave: %v\n\twant: %v", data[key], value)
	}
}

func TestExtenderError(t *testing.T) {
	message := "message"
	w := httptest.NewRecorder()
	_, err := gojsend.ExtendWriter(w).Error(message).Send()
	if err != nil {
		t.Errorf("ExtenderError\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("ExtenderError\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[gojsend.FieldStatus] != gojsend.StatusError {
		t.Errorf("ExtenderError\n\thave: %v\n\twant: %v", body[gojsend.FieldStatus], gojsend.StatusError)
	}
	if body[gojsend.FieldMessage] != message {
		t.Errorf("ExtenderError\n\thave: %v\n\twant: %v", body[gojsend.FieldMessage], message)
	}
}

func TestExtenderData(t *testing.T) {
	key := "foo"
	value := "bar"
	w := httptest.NewRecorder()
	_, err := gojsend.ExtendWriter(w).Data(map[string]interface{}{key: value}).Send()
	if err != nil {
		t.Errorf("ExtenderData\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("ExtenderData\n\thave: %v\n\twant: %v", err, nil)
	}
	data, ok := body[gojsend.FieldData].(map[string]interface{})
	if !ok {
		t.Errorf("ExtenderData\n\thave: %T\n\twant: %v", data, "map[string]interface{}")
	}
	if data[key] != value {
		t.Errorf("ExtenderData\n\thave: %v\n\twant: %v", data[key], value)
	}
}

func TestExtenderCode(t *testing.T) {
	code := 0
	w := httptest.NewRecorder()
	_, err := gojsend.ExtendWriter(w).Code(code).Send()
	if err != nil {
		t.Errorf("ExtenderCode\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("ExtenderCode\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[gojsend.FieldCode] != float64(code) {
		t.Errorf("ExtenderCode\n\thave: %v\n\twant: %v", body[gojsend.FieldCode], code)
	}
}

func TestExtenderJSONEncoder(t *testing.T) {
	w := httptest.NewRecorder()
	_, err := gojsend.ExtendWriter(w).JSONEncoder(json.Marshal).Send()
	if err != nil {
		t.Errorf("ExtenderJSONEncoder\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("ExtenderJSONEncoder\n\thave: %v\n\twant: %v", err, nil)
	}
}

func TestExtenderSend(t *testing.T) {
	w := httptest.NewRecorder()
	_, err := gojsend.ExtendWriter(w).Send()
	if err != nil {
		t.Errorf("ExtenderSend\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("ExtenderSend\n\thave: %v\n\twant: %v", err, nil)
	}
}
