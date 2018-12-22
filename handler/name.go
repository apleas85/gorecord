package handler

import (
	"net/http"

	"github.com/go/gorecords/service"
)

func Name(w http.ResponseWriter, r *http.Request) {
	service.SortByLastName()
	Response(w, r)
}
