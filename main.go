package main

import (
	"fmt"
	"math"
)
// function with one param, no return value
func sayGreeting(name string) {
	fmt.Printf("Good Morning %v \n", name)
}

// function with one param, no return value
func sayBye(name string)  {
	fmt.Printf("Good Bye %v \n", name)
}

// function with one param, one return value
func circleArea(radius float32) float32 {
	area := math.Pi * radius *radius
	return area
}

// function with two params string and function, no return value
func cycleName(listOfName []string, callBack func(string)) {
	for  _, value := range listOfName { // _ will avoid the index value.
		callBack(value)
	}
}

func main() {
	names :=[]string{"Abi", "Daniel", "Abishek"}
	sayGreeting("Daniel")
	sayGreeting("Abi")
	sayBye("Daniel")
	sayBye("Abi")
	cycleName(names, sayGreeting)
	cycleName(names, sayBye)
	areaOfCircle := circleArea(25.85)
	fmt.Printf("The area of the circle is %v", areaOfCircle) 
}