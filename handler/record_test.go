package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordHandlerFail(t *testing.T) {
	data := "shouldFailOnthis"
	req, err := http.NewRequest("Post", "/records", strings.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Record)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, 500)

}
