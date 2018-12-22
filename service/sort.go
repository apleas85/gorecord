package service

import (
	"fmt"
	"sort"
	"time"
)

//Sorts Records slice of Person structs in order by Gender (female first) then last name ascending
func SortByGender() {
	sort.Slice(Record.Records, func(i, j int) bool {
		if Record.Records[i].Gender == "Female" && Record.Records[j].Gender == "Female" {
			return Record.Records[i].LastName < Record.Records[j].LastName
		} else if Record.Records[i].Gender == "Male" && Record.Records[j].Gender == "Male" {
			return Record.Records[i].LastName < Record.Records[j].LastName
		} else {
			return Record.Records[i].Gender < Record.Records[j].Gender
		}
	})

}

//Sorts Records slice of Person structs by last name descending
func SortByLastName() {
	sort.Slice(Record.Records, func(i, j int) bool { return Record.Records[i].LastName > Record.Records[j].LastName })
}

//Sorts Records slice of Person structs by birthday ascending
func SortByBirthday() {
	sort.Slice(Record.Records, func(i, j int) bool {
		day1, err := time.Parse("01/02/2006", Record.Records[i].Birthday)
		if err != nil {
			fmt.Printf("error parsing date %v  ", err)
		}
		day2, err := time.Parse("01/02/2006", Record.Records[j].Birthday)
		if err != nil {
			fmt.Printf("error parsing date %v  ", err)
		}
		return day1.Before(day2)
	})
}
