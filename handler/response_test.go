package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseHandler(t *testing.T) {
	data := "shouldFailOnthis"
	req, err := http.NewRequest("Post", "/records", strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Response)
	handler.ServeHTTP(rr, req)
	// Check the status code is what we expect.
	assert.Equal(t, rr.Code, 200)

}
