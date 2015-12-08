/*Write a function with one variadic parameter that finds the greatest number in a list of numbers.*/
package main

import "log"

func main() {
	arr := []int{12,19,87,21,51}
	log.Println(maxNum(arr...))
}

func maxNum(args ...int) int {
	max := args[0]
	
	for i := 0; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
	}
	return max
}

