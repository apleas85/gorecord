package handler

import (
	"fmt"
	"net/http"
)

func BadPathHandler(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Page does noe exist:  %s", r.URL)
	http.Error(w, response, 404)
}
