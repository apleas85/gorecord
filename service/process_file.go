package service

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Records struct {
	Records []Person `json:"Records"`
}

type Person struct {
	LastName      string `json:"LastName"`
	FirstName     string `json:"FirstName"`
	Gender        string `json:"Gender"`
	FavoriteColor string `json:"FavoriteColor"`
	Birthday      string `json:"Birthday"`
}

var (
	Record Records
)

//takes a list of strings representing a file
//returns true if success or an error
func ProcessFiles(fileList []string) (bool, error) {
	for _, f := range fileList {
		absPath, _ := filepath.Abs(f)
		file, err := os.Open(absPath)
		if err != nil {
			return false, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			_, err := FormatLine(scanner.Text())
			if err != nil {
				return false, err
			}

		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	return true, nil
}

//splits and seperates file string and appends as a Person in a Records Struct
//returns true and nil or an error and false on fail
func FormatLine(value string) (bool, error) {
	var record []string
	if strings.Contains(value, " | ") {
		record = strings.Split(value, " | ")
	} else if strings.Contains(value, ", ") {
		record = strings.Split(value, ", ")
	} else if strings.Contains(value, " ") {
		record = strings.Split(value, " ")
	} else {
		return false, errors.New("Please enter a string of files seperated by | , or a space")
	}
	if len(record) != 5 {
		return false, fmt.Errorf("You must enter 5 values you entered %d", len(record))
	}
	Record.Records = append(Record.Records, Person{record[0], record[1], record[2], record[3], record[4]})

	return true, nil
}
