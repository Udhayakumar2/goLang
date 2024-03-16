package main

import (
	"fmt"
)

func main() {
	age := 45
	fmt.Println(age <= 50) // true
	fmt.Println(age >= 50) // false
	fmt.Println(age == 45) // true
	fmt.Println(age != 50) // true

	// if else condition
	if age < 30 {
		fmt.Println("Age is less than 30")
	} else {
		fmt.Println("Age is greater than 30")
	}

	// if else if else condtions
	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age should not be above 40")
	}
}