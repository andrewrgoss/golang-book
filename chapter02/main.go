package main

/*The first statement in a Go source file must be package name, where name is the package's default name for imports (all files in a package must use the same name). Executable commands must always use package main.

There are two types of Go programs: executables and libraries. Executable applications are the kinds of programs that we can run directly from the terminal. (in Windows they end with .exe) Libraries are collections of code that we package together so that we can use them in other programs.*/

import "fmt" // fmt = shorthand for 'format'

func main() {
	fmt.Println("Hello world") // sends whatever you give it to standard output (in the terminal you're working in)
}

/*Functions are the building blocks of a Go program. They have inputs, outputs and a series of steps called statements which are executed in order. All functions start with the keyword func followed by the name of the function (main in this case), a list of zero or more “parameters” surrounded by parentheses, an optional return type and a “body” which is surrounded by curly braces. This function has no parameters, doesn't return anything and has only one statement. The name main is special because it's the function that gets called when you execute the program.*/