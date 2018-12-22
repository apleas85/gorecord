package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadPathHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/records/badpath", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(BadPathHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, 404)

}
