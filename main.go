package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"Plan Dosa":   20,
		"Nice Dosa":   25,
		"Kal Dosa":    20,
		"Masala Dosa": 40,
	}

	fmt.Println(menu)
	fmt.Println(menu["Nice Dosa"])
	menu["Egg Dosa"] = 45

	// looping the maps
	for key, value := range menu {
		fmt.Println(key, "- Rs:", value,"/-")
	}


}
