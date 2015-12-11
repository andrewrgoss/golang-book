package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ { // Goroutines are lightweight and we can easily create thousands of them.
		go f(i) // f prints out the numbers from 0 to 10, waiting between 0 and 250 ms after each one. The goroutines should now run simultaneously.
	}
	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}