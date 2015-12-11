package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan string = make(chan string) // Channels provide a way for two goroutines to communicate with one another and synchronize their execution.
	
	go pinger(c)
	go ponger(c)
	go printer(c)
	
	var input string
	fmt.Scanln(&input)
}

func pinger(c chan string) { // A channel type is represented with the keyword chan followed by the type of the things that are passed on the channel (in this case we are passing strings).
	for i := 0; ; i++ {
		c <- "ping" // The <- (left arrow) operator is used to send and receive messages on the channel. c <- "ping" means send "ping".
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) { // channel direction specified, remove the <- and it is bi-directional
	for {
		msg := <- c // This means receive a message and store it in msg.
		fmt.Println(msg)
		time.Sleep(time.Second * 1) // when pinger attempts to send a message on the channel it will wait until printer is ready to receive the message (this is known as blocking)
	}
}