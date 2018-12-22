package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortByBirthday(t *testing.T) {
	fmt.Printf("Sort by birthday:    %v ", Record)

	Record.Records = []Person{Person{"Pleas", "Andrew", "Male", "FavoriteColorTest", "07/18/1985"}, Person{"Jefferson", "Tammy", "Female", "FavoriteColorTest", "04/20/1942"},
		Person{"Smith", "Jim", "Male", "FavoriteColorTest", "05/15/1999"}, Person{"Peters", "Terry", "Female", "FavoriteColorTest", "05/15/2017"}}
	SortByBirthday()
	assert.Equal(t, 4, len(Record.Records), "they should be equal")
	assert.Equal(t, "Tammy", Record.Records[0].FirstName, "they should be equal")

}

func TestSortByLastName(t *testing.T) {
	fmt.Printf("Sort by birthday:    %v ", Record)
	Record.Records = []Person{Person{"Pleas", "Andrew", "Male", "FavoriteColorTest", "07/18/1985"}, Person{"Jefferson", "Tammy", "Female", "FavoriteColorTest", "04/20/1942"},
		Person{"Smith", "Jim", "Male", "FavoriteColorTest", "05/15/1999"}, Person{"Peters", "Terry", "Female", "FavoriteColorTest", "05/15/2017"}}
	SortByLastName()
	assert.Equal(t, 4, len(Record.Records), "they should be equal")
	assert.Equal(t, "Smith", Record.Records[0].LastName, "they should be equal")

}

func TestSortByGender(t *testing.T) {
	Record.Records = []Person{Person{"Pleas", "Andrew", "Male", "FavoriteColorTest", "07/18/1985"}, Person{"Jefferson", "Tammy", "Female", "FavoriteColorTest", "04/20/1942"},
		Person{"Smith", "Jim", "Male", "FavoriteColorTest", "05/15/1999"}, Person{"Peters", "Terry", "Female", "FavoriteColorTest", "05/15/2017"}}
	SortByGender()
	fmt.Printf("Sort by gender:    %v ", Record)
	assert.Equal(t, 4, len(Record.Records), "they should be equal")
	assert.Equal(t, "Female", Record.Records[0].Gender, "they should be equal")
	assert.Equal(t, "Female", Record.Records[1].Gender, "they should be equal")
	assert.Equal(t, "Male", Record.Records[2].Gender, "they should be equal")
	assert.Equal(t, "Male", Record.Records[3].Gender, "they should be equal")

}
