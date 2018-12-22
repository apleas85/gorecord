package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatLineComma(t *testing.T) {
	Record.Records = nil
	Test, err := FormatLine("LastNameTest, FirstNameTest, GenderTest, FavoriteColorTest, DateOfBirthTest")
	assert.Equal(t, len(Record.Records), 1, "they should be equal")
	assert.True(t, Test, "Should be true")
	assert.Nil(t, err, "Error should be nil")
}

func TestFormatLinePipe(t *testing.T) {
	Record.Records = nil
	Test, err := FormatLine("LastNameTest | FirstNameTest | GenderTest | FavoriteColorTest | DateOfBirthTest")
	assert.Equal(t, len(Record.Records), 1, "they should be equal")
	assert.True(t, Test, "Should be true")
	assert.Nil(t, err, "Error should be nil")
}

func TestFormatLineSpace(t *testing.T) {
	Record.Records = nil
	Test, err := FormatLine("LastNameTest FirstNameTest GenderTest FavoriteColorTest DateOfBirthTest")
	assert.Equal(t, len(Record.Records), 1, "they should be equal")
	assert.True(t, Test, "Should be true")
	assert.Nil(t, err, "Error should be nil")
}

func TestFormatLineError(t *testing.T) {
	Record.Records = nil
	Test, err := FormatLine("ThisWillFail")
	assert.NotNil(t, err)
	assert.False(t, Test, "Should be false")
}

func TestFormatLineError2(t *testing.T) {
	Record.Records = nil
	Test, err := FormatLine("ThisW illFail")
	assert.NotNil(t, err)
	assert.False(t, Test, "Should be false")
	assert.Equal(t, err.Error(), "You must enter 5 values you entered 2")
}

func TestProcessFile(t *testing.T) {
	Record.Records = nil
	var files = []string{"testFiles/spaceRecordsTest.txt"}
	test, err := ProcessFiles(files)
	assert.Equal(t, len(Record.Records), 4, "they should be equal")
	assert.True(t, test, "Should be true")
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, Record.Records[0].FirstName, "Rio", "Error should be nil")
}
func TestProcessFileFail(t *testing.T) {
	Record.Records = nil
	var files = []string{"testFiles/space.txt"}
	test, err := ProcessFiles(files)
	assert.False(t, test, "Should be true")
	assert.NotNil(t, err, "Error should not be nill")
}
