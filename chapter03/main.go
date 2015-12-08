package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1 + 1 =", 1.0 + 1.0) //.0 tells Go that this is a floating point number instead of an integer. Go is able to infer this data type.
}