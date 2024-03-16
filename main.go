package main

import "fmt"

func main() {
	age := 25
	name := "Daniel"

	// Print
	fmt.Print("Hello") // This will print the text
	fmt.Print("World! \n") // \n is used to print the text in the next line
	fmt.Print("new line \n") // The "" is used to refer the strings and ''  refer to the characters
	
	// Println
	fmt.Println("Hello Daniel!") // Println always print in the new line.
	fmt.Println("My age is ", age, " and my name is ", name)

	// Printf (formated string) %_ = format specifier
	fmt.Printf("My age is %v and my name is %v \n", age, name) //%v assign the variables
	fmt.Printf("My age is %q and my name is %q \n", age, name) // %q print the string values inside a quotes.
	fmt.Printf("My age is %T and my name is %T \n", age, name) // %T print the data type of the variable.
	fmt.Printf("Your score is %f \n", 25.05) // %f print float value like 25.050000
	fmt.Printf("Your score is %0.1f \n", 25.05) // %0.1f print rounded float value like 25.1

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println(str)

  
}