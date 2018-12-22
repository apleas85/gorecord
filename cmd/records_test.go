package cmd

import (
	"testing"

	"github.com/go/gorecords/service"
	"github.com/stretchr/testify/assert"
)

func TestRecorsVerifyFiles(t *testing.T) {
	ok, err := validateFiles("files", "file1.txt,file2.txt,file3.txt")
	assert.True(t, ok)
	assert.Nil(t, err)

}
func TestReadRecordsFail1(t *testing.T) {
	files = "file1.txt,file2.txt,file3.txt"
	val, err := readRecordsHelper()
	assert.Equal(t, "", val)
	assert.NotNil(t, err)

}
func TestReadRecordsFail2(t *testing.T) {
	files = "file1.txt file2.txt file3.txt"
	val, err := readRecordsHelper()
	assert.Equal(t, "", val)
	assert.NotNil(t, err)
	assert.Equal(t, "Error validating file string:  Please provide three files seperated by commas", err.Error())

}

func TestReadRecordsFail3(t *testing.T) {
	files = ""
	val, err := readRecordsHelper()
	assert.Equal(t, "", val)
	assert.NotNil(t, err)
	assert.Equal(t, "Error validating file string:  Please provide a valid file string", err.Error())

}

func TestReadRecords(t *testing.T) {
	files = "../spaceRecords.txt,../commaRecords.txt,../pipeRecords.txt"
	val, err := readRecordsHelper()
	assert.Equal(t, len(service.Record.Records), 10)
	assert.NotNil(t, val)
	assert.Nil(t, err)
	byBirthday = true
	sortRecords()
	assert.Equal(t, service.Record.Records[0].FirstName, "Tod")
	byBirthday = false
	byLastName = true
	sortRecords()
	assert.Equal(t, service.Record.Records[0].LastName, "Smith")
	byLastName = false
	byGender = true
	sortRecords()
	assert.Equal(t, service.Record.Records[0].Gender, "Female")
	byGender = false

}
