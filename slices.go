package main

import (
	"fmt"
	)

func main() {
	parents := []string{"Carol", "Mike"}
	kids := []string{"Marcia", "Jan", "Cindy", "Greg", "Peter", "Bobby"}

	family := append(parents, kids...)

	fmt.Println(family)

	// Fix the following bugs
	extras := make([]string, 1, 1)
	extras[0] = "Alice"
	fmt.Println(extras)
}