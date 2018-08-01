package main

import "fmt"

func main() {

	a := "Hello"
	b := " World"

	x := 1
	y := 2

	c := a + b
	z := x + y

	fmt.Printf("c is %v\n", c)
	s := fmt.Sprintf("Z is %v", z)
	fmt.Println(s)
}
