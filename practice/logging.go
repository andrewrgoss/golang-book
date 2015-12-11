package main

import (
	"log"
	"os"
)

func main() {

	filename := "log.txt"

	file, err := os.Create(filename) // create new log.txt file
    file.Close()

	// make sure said log.txt exists...
	logFile, err := os.OpenFile("log.txt", os.O_WRONLY, 0666)

	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	// direct all log messages to log.txt
	log.SetOutput(logFile)

	log.Println("First log message!")
    log.Println("Second log message!")
    log.Println("Third log message!")
}
