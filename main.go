package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	// In this session you will learn about standard library(default libraries) in this exercise we are going to use strings and sort library.
	greeting := "Hello there friends!"

	// strings
	fmt.Println(strings.Contains(greeting, "Hello")) // output: true
	fmt.Println(strings.Contains(greeting, "Hello!")) // output: false
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi")) // output: Hi there friends!
	fmt.Println(strings.ToUpper(greeting)) // output: HELLO THERE FRIENDS!
	fmt.Println(strings.Index(greeting, "fri")) // output: 12
	fmt.Println(strings.Split(greeting, " ")) // output: [Hello there friends!]

	// sort
	ages :=[]int{20,20,24,15,28,23,25,18}
	fmt.Println(ages) // output: [20 20 24 15 28 23 25 18]
	sort.Ints(ages)
	index := sort.SearchInts(ages, 20)
	fmt.Println(index) // output: 2
	fmt.Println(ages) // output: [15 18 20 20 23 24 25 28]
	
}