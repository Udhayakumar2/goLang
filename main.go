package main
 // If the package name are same, They can share all the function, variables, etc with in that package.
 
import (
	"fmt"
)
func main() {
	sayGreeting("Daniel")
	sayGreeting("Abi")
	sayBye("Daniel")
	sayBye("Abi")
	cycleName(names, sayGreeting)
	cycleName(names, sayBye)
	areaOfCircle := circleArea(25.85)
	fmt.Printf("The area of the circle is %v", areaOfCircle) 
}