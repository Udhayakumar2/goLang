package main

import ("fmt")

func main() {
	var ages [3]int = [3]int{20, 22, 23} // Declaring Array [array size]datatype
	name := [4]string{"Daniel", "Abishek", "Dani", "Abi"} // Declaring Array in short hand method.
	fmt.Println(ages, len(ages))
	fmt.Println(name, len(name))

	// slices (use arrays under the hood)
	var scores = []int{100, 20,50} // Declaring array without mentioning the size
	scores[2] = 51 // updating the value by index value
	scores = append(scores, 85) // Appending the new value into the array

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := name[1:3]
	fmt.Println(rangeOne)

}