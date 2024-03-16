package main

import (
	"fmt"
)

func assignValueInMap(mapData map[string]float64, key string, value float64) {
	mapData[key] = value
}

func assignValueInString(getValue string) {
	getValue = "Abishek"
}

func main() {
	menu := map[string]float64{
		"Plan Dosa":   20,
		"Nice Dosa":   25,
		"Kal Dosa":    20,
		"Masala Dosa": 40,
	}
	name := "Daniel"

	fmt.Println(menu)
	fmt.Println(menu["Nice Dosa"])
	assignValueInMap(menu, "Egg Dosa", 45)
	assignValueInString(name)
	// looping the maps
	for key, value := range menu {
		fmt.Println(key, "- Rs:", value,"/-")
	}

	fmt.Println(name) // the name will not change because it pass the copy of the parent value, Output: Daniel
	fmt.Println(menu) // The menu will change because the map pass it's pointer value(memory address) Output: map[Egg Dosa:45 Kal Dosa:20 Masala Dosa:40 Nice Dosa:25 Plan Dosa:20]



}
