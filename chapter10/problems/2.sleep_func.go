//Write your own Sleep function using time.After.
package main

import (
	"fmt"
	"time"
)

func main() {
	// bi-directional channel
	c := make(chan string)
	
	go sleep(c)
		
	var input string
	fmt.Scanln(&input)
}

func sleep(c chan string) {
	for {
		select {
			case msg1 := <-c:
				fmt.Println("Message 1", msg1)
			case <-time.After(time.Second * 2):
				fmt.Println("Sleep...")
		}
	}
}