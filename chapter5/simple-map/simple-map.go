package main

import (
	"fmt"
)

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Li"])
	val, ok := elements["Un"]
	fmt.Println(val, ok)

	if val, ok := elements["Un"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("The element \"Un\" was not found in the dictionary.")
	}

	elementsAlt := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elementsAlt)

	descriptiveElements := map[string](map[string]string){
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
	}

	fmt.Println(descriptiveElements)

	if description, ok := descriptiveElements["H"]; ok {
		fmt.Println(description["name"], description["state"])
	}
}
