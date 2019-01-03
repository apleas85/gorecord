package server

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go/gorecords/handler"
	"github.com/go/gorecords/service"
	"github.com/gorilla/mux"
)

const (
	port = 9010
)

//loads file into Record struct
func loadFiles(files string) error {
	fileList := strings.Split(files, ",")
	_, err := service.ProcessFiles(fileList)
	if err != nil {
		fmt.Println(err)
		response := fmt.Sprintf("Error processing files:  %v", err)
		return errors.New(response)
	}
	return nil
}

//starts server on port 9010
func Start(files string) error {
	r := mux.NewRouter()
	r.HandleFunc("/records/gender", handler.Gender)
	r.HandleFunc("/records/birthday", handler.Birthday)
	r.HandleFunc("/records/name", handler.Name)
	r.HandleFunc("/records", handler.Record).Methods("POST")

	if err := loadFiles(files); err != nil {
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
