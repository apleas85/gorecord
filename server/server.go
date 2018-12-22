package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go/gorecords/handler"
	"github.com/go/gorecords/service"
	"github.com/gorilla/mux"
)

const (
	port = 9010
)

var files = []string{"spaceRecords.txt", "commaRecords.txt", "pipeRecords.txt"}

//loads file into Record struct
func loadFiles() error {
	_, err := service.ProcessFiles(files)
	if err != nil {
		fmt.Println(err)
		response := fmt.Sprintf("Error processing files:  %v", err)
		return errors.New(response)
	}
	return nil
}

//starting server on port 9010
func Start() error {
	r := mux.NewRouter()
	r.HandleFunc("/records/gender", handler.Gender)
	r.HandleFunc("/records/birthday", handler.Birthday)
	r.HandleFunc("/records/name", handler.Name)
	r.HandleFunc("/records", handler.Record).Methods("POST")

	if err := loadFiles(); err != nil {
		fmt.Println("Error loading files ")
		return err
	}
	fmt.Printf(fmt.Sprintf("Starting server at port %d  \n", port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
		fmt.Printf("Error starting server %v", err.Error())
		return err
	}
	return nil
}
