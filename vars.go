package main

import "fmt"

func main() {
	// Declare the following variables with zero values:
	// name: command, type string
	var command string

	// name: valid, type bool
	var valid bool

	// print the value of those variables
	fmt.Printf("Command = %q\n", command)
	fmt.Printf("Valid = %v\n", valid)

	// declare and initialize the following variables:
	// name: foo, type: int, value: 5
	foo := 5

	// name bar, type: bool, value: true
	bar := true

	// print the value of those variables
	fmt.Printf("foo = %v\n", foo)
	fmt.Printf("bar = %v\n", bar)
}