package gojsend_test

import (
	"net/http/httptest"
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

func TestNewWriter(t *testing.T) {
	w := httptest.NewRecorder()
	j := gojsend.NewWriter(w)
	j, ok := j.(gojsend.JSendWriter)
	if !ok {
		t.Errorf("NewWriter\n\thave: %T\n\twant: %v", j, "JSendWriter")
	}
}
