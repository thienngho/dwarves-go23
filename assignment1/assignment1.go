package main

import (
	"fmt"
	"os"
	"strings"
)

func ValidateCountryCode(countryCode string) bool {
	switch countryCode {
	case "US":
		return true
	case "VN":
		return true
	}
	return false
}

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("Usage: go run assignment_1.go <first name> <last name> [middle name(s)] <country code>")
		os.Exit(1)
	}

	firstName := args[0]
	lastName := args[1]
	countryCode := args[len(args)-1]

	if !ValidateCountryCode(countryCode) {
		fmt.Println("Missing country code")
		os.Exit(1)
	}

	middleNames := args[2 : len(args)-1]
	middleName := strings.Join(middleNames, " ")

	var reorderedName string

	switch countryCode {
	case "US":
		reorderedName = fmt.Sprintf("%s %s %s", firstName, middleName, lastName)
	case "VN":
		reorderedName = fmt.Sprintf("%s %s %s", lastName, middleName, firstName)
	default:
		reorderedName = fmt.Sprintf("%s %s %s", firstName, middleName, lastName)
	}

	fmt.Println(reorderedName)
}
