/*sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?*/

package main

import "log"

func main() {
	mySlice := []int{5,4,5,6}
	log.Println(sum(mySlice))
}

func sum(numSlice []int) (sum int) {
	sum = 0
	for v := range numSlice {
		sum += numSlice[v]
	}
	return
}


// func main() {
// 	nums := []int{5,4,5,6}
// 	log.Println(sum(nums...))
// }
// 
// func sum(args ...int) int {
// 	total := 0
// 	for _, v := range args {
// 		total += v
// 	}
// 	return total
// }