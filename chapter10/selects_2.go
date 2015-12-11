package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

go func() {
	for {
		c1 <- "from 1"
		time.Sleep(time.Second * 2)
	}
}()

go func() {
	for {
		c2 <- "from 2"
		time.Sleep(time.Second * 3)
	}
}()

go func() {
	for {
		select { // select picks the first channel that is ready and receives from it (or sends to it). If more than one of the channels are ready then it randomly picks which one to receive from. If none of the channels are ready, the statement blocks until one becomes available.
			case msg1 := <- c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <- c2:
				fmt.Println("Message 2", msg2)
			case <- time.After(time.Second): // this creates a channel and after the given duration will send the current time on it. (we weren't interested in the time so we didn't store it in a variable)
				fmt.Println("timeout")
			default:
				fmt.Println("nothing ready")
			}
		}
	}()
	
	var input string
	fmt.Scanln(&input)
}