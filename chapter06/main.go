//https://www.golang-book.com/books/intro/6

package main

import "fmt"

func main() {
	// var x [5]int
	// x[4] = 100
	// fmt.Println(x)

	// var x [5]float64
	// x[0] = 98
	// x[1] = 93
	// x[2] = 77
	// x[3] = 82
	// x[4] = 83

	// x := [5]float64{98, 93, 77, 82, 83}

	x := [5]float64{
		98,
		93,
		77,
		82,
		83,
	}

	// var total float64 = 0
	// for i := 0; i < len(x); i++ {
	// 	total += x[i]
	// }
	// fmt.Println(total / float64(len(x)))

	var total float64 = 0
	for _, value := range x { //A single _ (underscore) is used to tell the compiler that we don't need this. (In this case we don't need the iterator (i) variable)
		total += value
	}
	fmt.Println(total / float64(len(x)))

	/*A slice is a segment of an array. Like arrays slices are indexable and have a length. Unlike arrays this length is allowed to change.*/

	//var y []float64 //The only difference between this and an array is the missing length between the brackets. In this case y has been created with a length of 0.

	//y := make([]float64, 5, 10) //5 represents the slice, 10 represents the capacity of the underlying array which the slice points to.

	arr := [5]float64{1, 2, 3, 4, 5}
	y := arr[:] //or arr[0:5] or arr[0:] or arr[0:len(arr)] or arr[:5]
	fmt.Println(y)

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("Append slice example:", slice1, slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println("Copy slice example", slice3, slice4)

	/*A map is an unordered collection of key-value pairs. Also known as an associative array, a hash table or a dictionary, maps are used to look up a value by its associated key. Here's an example of a map in Go:
	 */

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
	}

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
