package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFiles(t *testing.T) {
	files = []string{"IdoneExist.txt", "iwillfail.txt", "badfile.txt"}
	err := loadFiles()
	assert.NotNil(t, err)
	assert.Equal(t, port, 9010)

}

func TestStart(t *testing.T) {
	files = []string{"IdoneExist.txt", "iwillfail.txt", "badfile.txt"}
	err := Start()
	assert.Equal(t, port, 9010)
	assert.NotNil(t, err)
}
