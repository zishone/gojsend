package gojsend_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/zishone/gojsend"
)

func TestDecoratorJSONEncoder(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/login", strings.NewReader(""))
	h := func(j gojsend.JSendWriter, r *http.Request) {
		j.Success(map[string]interface{}{"foo": "bar"}).
			Send()
	}
	j := gojsend.NewDecorator().
		JSONEncoder(json.Marshal).
		Decorate(h)
	j(w, req)
	body := make(map[string]interface{})
	err := json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("DecoratorJSONEncoder\n\thave: %v\n\twant: %v", err, nil)
	}
}

func TestDecoratorDecorate(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/login", strings.NewReader(""))
	h := func(j gojsend.JSendWriter, r *http.Request) {
		j.Success(map[string]interface{}{"foo": "bar"}).
			Send()
	}
	j := gojsend.NewDecorator().
		Decorate(h)
	j(w, req)
	body := make(map[string]interface{})
	err := json.NewDecoder(w.Result().Body).Decode(&body)
	if err != nil {
		t.Errorf("DecoratorDecorate\n\thave: %v\n\twant: %v", err, nil)
	}
}
