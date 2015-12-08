//What's the value of the expression (true && false) || (false && true) || !(false && false)?
package main

import "fmt"

func main() {
	fmt.Println(evaluate()) // true
}

func evaluate()bool {
	return (true && false) || (false && true) || !(false && false)
}

// true && false = false
// false && true = false
// false && false = false
// Therefore, false || false || !false = true