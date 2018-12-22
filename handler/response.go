package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go/gorecords/service"
)

func Response(w http.ResponseWriter, r *http.Request) {
	jsonRecords, err := json.MarshalIndent(service.Record, "", "  ")
	if err != nil {
		fmt.Println(err)
		response := fmt.Sprintf("Could not parse Json:  %v", err)
		http.Error(w, response, 500)
	}
	fmt.Fprintln(w, string(jsonRecords))
}
