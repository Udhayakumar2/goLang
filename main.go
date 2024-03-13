package main

import "fmt"

func main() {
	var title string = "Hello World"
	var titleOne = "Good Morning"
	var titleTwo string

	fmt.Println(title, titleOne, titleTwo)

	title = "Daniel"
	titleTwo = "Hello World"
	fmt.Println(title, titleOne, titleTwo)

	varShortHand := "This is the short hand for var" // This will be work only inside the function scope not on outer.

	fmt.Println(title, titleOne, titleTwo, varShortHand)

	// data type
	var name string = "Udhaya"
	var age int = 25
	// bit&memory

	// int
	var integer8 int8 = 127                    // Range -128 to 127.
	var integer16 int16 = 32767                // Range -32768 to 32767.
	var integer32 int32 = 2147483647           // int and int32 are same. Range -2147483648 to 2147483647.
	var integer64 int64 = -9223372036854775808 // Range -9223372036854775808 to 9223372036854775807
	fmt.Println("Integers: ",name, age, integer8, integer16, integer32, integer64)
	
	// uint
	var uInt8 uint8 = 255 // Range 0 to 255
	var uInt16 uint16 = 65535 // Range 0 to 65535
	var uInt32 uint32 = 4294967295 // uint and uint35 are same. Range 0 to 4294967295
	var uInt64 uint64 = 18446744073709551615 // Range 0 to 18446744073709551615
	fmt.Println("Positive Integers: ", uInt8, uInt16, uInt32, uInt64)

	// float
	var floatInt32 float32 = 2.5
	var floatInt64 float64 = 549865184615165198.151 // float64 is the default one.
	floatInt := 2.5 // hover the variable name to see the default datatype
	fmt.Println(floatInt32, floatInt64, floatInt)

}
