package handler

import (
	"net/http"

	"github.com/go/gorecords/service"
)

// Ping reponds with a 200 ok message if server is available

func Gender(w http.ResponseWriter, r *http.Request) {
	service.SortByGender()
	Response(w, r)
}
