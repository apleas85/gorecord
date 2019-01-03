package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFiles(t *testing.T) {
	err := loadFiles("IdoneExist.txt,iwillfail.txt,badfile.txt")
	assert.NotNil(t, err)
	assert.Equal(t, port, 9010)

}

func TestStart(t *testing.T) {
	err := Start("IdoneExist.txt,iwillfail.txt,badfile.txt")
	assert.Equal(t, port, 9010)
	assert.NotNil(t, err)
}
