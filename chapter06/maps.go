package main

import "fmt"

/*A map is an unordered collection of key-value pairs. Also known as an associative array, a hash table or a dictionary, maps are used to look up a value by its associated key. Here's an example of a map in Go:*/

func main() {

	// var x map[string]int
	// x["key"] = 10
	// fmt.Println(x)

	//Maps have to be initialized before they can be used

	z := make(map[string]int)
	z["key"] = 10
	fmt.Println(z["key"])

	a := make(map[int]int)
	a[1] = 10
	fmt.Println(a[1])

	// elements := make(map[string]string)
	// elements["H"] = "Hydrogen"
	// elements["He"] = "Helium"
	// elements["Li"] = "Lithium"
	// elements["Be"] = "Beryllium"
	// elements["B"] = "Boron"
	// elements["C"] = "Carbon"
	// elements["N"] = "Nitrogen"
	// elements["O"] = "Oxygen"
	// elements["F"] = "Fluorine"
	// elements["Ne"] = "Neon"
	
	//Here's a shorter way to create maps..

	// elements := map[string]string{
	// 	"H":  "Hydrogen",
	// 	"He": "Helium",
	// 	"Li": "Lithium",
	// 	"Be": "Beryllium",
	// 	"B":  "Boron",
	// 	"C":  "Carbon",
	// 	"N":  "Nitrogen",
	// 	"O":  "Oxygen",
	// 	"F":  "Fluorine",
	// 	"Ne": "Neon",
	// }

	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	} // We now have a map of strings to maps of strings to strings. The outer map is used as a lookup table based on the element's symbol, while the inner maps are used to store general information about the elements.

	fmt.Println(elements["Li"])

	// name, ok := elements["Un"]
	// fmt.Println(name, ok)

	//accessing an element of a map can return two values instead of just one. The first value is the result of the lookup, the second tells us whether or not the lookup was successful.
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
	
	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}