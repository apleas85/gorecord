package handler

import (
	"net/http"

	"github.com/go/gorecords/service"
)

// Ping reponds with a 200 ok message if server is available
func Birthday(w http.ResponseWriter, r *http.Request) {
	service.SortByBirthday()

	Response(w, r)

}
