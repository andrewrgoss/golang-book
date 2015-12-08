//We used the Println function defined in the fmt package. If we wanted to use the Exit function from the os package what would we need to do?

package main

import "os"

func main() {
	os.Exit(0) // Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, non-zero an error. The program terminates immediately; deferred functions are not run.
}