package main

import "log"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func main() {
	someOtherName := []float64{98, 93, 77, 82, 83}
	log.Println(average(someOtherName))
}
