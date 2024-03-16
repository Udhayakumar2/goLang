package main

import (
	"fmt"
	"strings"
)

// function with one param and return two string value
// This function get the full name and return the initial of the first two names
func getInitialLetter(name string) (string, string) {
	splitName := strings.Split(strings.ToUpper(name), " ")
	var initials []string
	for _, value := range splitName {
		initials = append(initials, value[:1]) 
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	firstInitial, lastInitial := getInitialLetter("Daniel Abishek")
	fmt.Printf("First name initial is %q, last name initial is %q \n", firstInitial, lastInitial)
	firstInitial, lastInitial = getInitialLetter("Daniel")
	fmt.Printf("First name initial is %q, last name initial is %q \n", firstInitial, lastInitial)

}