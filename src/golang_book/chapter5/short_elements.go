package main

import "fmt"

func main() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"Ne": "Neon",
		"F":  "Fluorine",
	}

	fmt.Println(elements["C"])
}
