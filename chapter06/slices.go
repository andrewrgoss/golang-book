package main

import "fmt"

/*A slice is a segment of an array. Like arrays slices are indexable and have a length. Unlike arrays this length is allowed to change.*/

func main() {
	
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

}