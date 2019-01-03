package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go/gorecords/service"
	"github.com/spf13/cobra"
)

var (
	//files      string
	byLastName bool
	byGender   bool
	byBirthday bool
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Reads records from file",
	Run:   readRecords,
}

func init() {
	//	readCmd.Flags().StringVar(&files, "files", "spaceRecords.txt,commaRecords.txt,pipeRecords.txt", "Specify 3 files seperated by ','")
	readCmd.Flags().BoolVar(&byGender, "byGender", false, "When activated will sort by gender")
	readCmd.Flags().BoolVar(&byLastName, "byLastName", false, "When activated will sort by last name")
	readCmd.Flags().BoolVar(&byBirthday, "byBirthday", false, "When activated will sort by birthday")

}

func readRecords(cmd *cobra.Command, args []string) {
	_, err := readRecordsHelper(args...)
	if err != nil {
		response := fmt.Sprintf("Error processing files:  %v", err)
		fmt.Println(response)
		os.Exit(0)
	}
}
func readRecordsHelper(args ...string) (string, error) {
	if ok, err := validateFiles("file", files); !ok {
		response := fmt.Sprintf("Error validating file string:  %v", err)
		return "", errors.New(response)
	}

	fileList := strings.Split(files, ",")

	_, err := service.ProcessFiles(fileList)
	if err != nil {
		response := fmt.Sprintf("Error processing files:  %v", err)
		return "", errors.New(response)
		//os.Exit(0)
	}
	sortRecords()

	jsonRecords, err := json.MarshalIndent(service.Record, "", "  ")
	if err != nil {
		response := fmt.Sprintf("Error parsing JSON:  %v", err)
		return "", errors.New(response)
	}

	fmt.Println(string(jsonRecords))
	return string(jsonRecords), nil

}

func sortRecords() {
	if byGender {
		fmt.Println("Sorting by gender ")
		service.SortByGender()
	} else if byLastName {
		fmt.Println("Sorting by lastName ")
		service.SortByLastName()
	} else if byBirthday {
		fmt.Println("Sorting by birthday ")
		service.SortByBirthday()
	} else {
		fmt.Println("You must set one sorting flag byGender, byBirthday or byLastName.")
		os.Exit(0)

	}

}

func validateFiles(name, value string) (bool, error) {
	if len(value) == 0 {
		message := fmt.Sprintf("Please provide a valid %s string", name)
		err := errors.New(message)
		return false, err
	}
	if !strings.Contains(value, ",") {
		message := fmt.Sprintf("Please provide three files seperated by commas")
		err := errors.New(message)
		return false, err
	}
	return true, nil
}
