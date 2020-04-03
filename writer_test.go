package gojsend_test

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/zishone/gojsend"
)

func TestWriterSet(t *testing.T) {
	key := "foo"
	value := "bar"
	w := httptest.NewRecorder()
	if _, err := gojsend.NewWriter(w).Set(key, value).Send(); err != nil {
		t.Errorf("WriterSet\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err := json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("WriterSet\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[key] != value {
		t.Errorf("WriterSet\n\thave: %v\n\twant: %v", body[key], value)
	}
}

func TestWriterSuccess(t *testing.T) {
	key := "foo"
	value := "bar"
	w := httptest.NewRecorder()
	_, err := gojsend.NewWriter(w).Success(map[string]interface{}{key: value}).Send()
	if err != nil {
		t.Errorf("WriterSuccess\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("WriterSuccess\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[gojsend.FieldStatus] != gojsend.StatusSuccess {
		t.Errorf("WriterSuccess\n\thave: %v\n\twant: %v", body[gojsend.FieldStatus], gojsend.StatusSuccess)
	}
	data, ok := body[gojsend.FieldData].(map[string]interface{})
	if !ok {
		t.Errorf("WriterSuccess\n\thave: %T\n\twant: %v", data, "map[string]interface{}")
	}
	if data[key] != value {
		t.Errorf("WriterSuccess\n\thave: %v\n\twant: %v", data[key], value)
	}
}

func TestWriterFail(t *testing.T) {
	key := "foo"
	value := "bar"
	w := httptest.NewRecorder()
	_, err := gojsend.NewWriter(w).Fail(map[string]interface{}{key: value}).Send()
	if err != nil {
		t.Errorf("WriterFail\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("WriterFail\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[gojsend.FieldStatus] != gojsend.StatusFail {
		t.Errorf("WriterFail\n\thave: %v\n\twant: %v", body[gojsend.FieldStatus], gojsend.StatusFail)
	}
	data, ok := body[gojsend.FieldData].(map[string]interface{})
	if !ok {
		t.Errorf("WriterFail\n\thave: %T\n\twant: %v", data, "map[string]interface{}")
	}
	if data[key] != value {
		t.Errorf("WriterFail\n\thave: %v\n\twant: %v", data[key], value)
	}
}

func TestWriterError(t *testing.T) {
	message := "message"
	w := httptest.NewRecorder()
	_, err := gojsend.NewWriter(w).Error(message).Send()
	if err != nil {
		t.Errorf("WriterError\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("WriterError\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[gojsend.FieldStatus] != gojsend.StatusError {
		t.Errorf("WriterError\n\thave: %v\n\twant: %v", body[gojsend.FieldStatus], gojsend.StatusError)
	}
	if body[gojsend.FieldMessage] != message {
		t.Errorf("WriterError\n\thave: %v\n\twant: %v", body[gojsend.FieldMessage], message)
	}
}

func TestWriterData(t *testing.T) {
	key := "foo"
	value := "bar"
	w := httptest.NewRecorder()
	_, err := gojsend.NewWriter(w).Data(map[string]interface{}{key: value}).Send()
	if err != nil {
		t.Errorf("WriterData\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("WriterData\n\thave: %v\n\twant: %v", err, nil)
	}
	data, ok := body[gojsend.FieldData].(map[string]interface{})
	if !ok {
		t.Errorf("WriterData\n\thave: %T\n\twant: %v", data, "map[string]interface{}")
	}
	if data[key] != value {
		t.Errorf("WriterData\n\thave: %v\n\twant: %v", data[key], value)
	}
}

func TestWriterCode(t *testing.T) {
	code := 0
	w := httptest.NewRecorder()
	_, err := gojsend.NewWriter(w).Code(code).Send()
	if err != nil {
		t.Errorf("WriterCode\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("WriterCode\n\thave: %v\n\twant: %v", err, nil)
	}
	if body[gojsend.FieldCode] != float64(code) {
		t.Errorf("WriterCode\n\thave: %v\n\twant: %v", body[gojsend.FieldCode], code)
	}
}

func TestWriterJSONEncoder(t *testing.T) {
	w := httptest.NewRecorder()
	_, err := gojsend.NewWriter(w).JSONEncoder(json.Marshal).Send()
	if err != nil {
		t.Errorf("WriterJSONEncoder\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("WriterJSONEncoder\n\thave: %v\n\twant: %v", err, nil)
	}
}

func TestWriterSend(t *testing.T) {
	w := httptest.NewRecorder()
	_, err := gojsend.NewWriter(w).Send()
	if err != nil {
		t.Errorf("WriterSend\n\thave: %v\n\twant: %v", err, nil)
	}
	body := make(map[string]interface{})
	err = json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("WriterSend\n\thave: %v\n\twant: %v", err, nil)
	}
}
