package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go/gorecords/service"
)

type test_struct struct {
	Test string
}

func Record(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response := fmt.Sprintf("Error Reading body files:  %v", err)
		http.Error(w, response, 500)
	}
	body := string(bodyBytes)

	if _, err := service.FormatLine(body); err != nil {
		response := fmt.Sprintf("Error:  %v", err)
		http.Error(w, response, 500)
	}

	fmt.Fprintln(w, "Successfully wrote to file!")

}
