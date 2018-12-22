package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNameHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/records/name", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Name)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, 200)

}
