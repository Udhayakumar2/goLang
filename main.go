package main

import (
	"fmt"
)

func main () {
	x := 0
	// The below loop look like for loop but it is behave like a while loop
	for x < 5 {
		fmt.Println("Value of x is: ", x)
		x++
	}

	// Traditional for loop
	ages := []int{20, 22, 24, 15, 28, 23, 25, 18}
	for i := 0; i < len(ages); i++ {
		fmt.Printf("Tranditional for loop index %d: %d \n", i,ages[i])
	}

	// Short hand for loop
	for index, value := range ages {
		fmt.Printf("Short hand for loop index %d: %d \n", index, value)
	}
}