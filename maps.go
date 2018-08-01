package main

import "fmt"

func main() {

	beatles := map[string]string{}

	beatles["John"] = "guitar"
	beatles["Paul"] = "bass"
	beatles["George"] = "guitar"
	beatles["Ringo"] = "drums"

	// Loop through the map and print out the key and the value

	for k, v := range beatles {
		fmt.Printf( "%s : %s\n", k, v)
	}

	// Look up the key `Bob` and detect that it wasn't found by printing `not found`
	if _, ok := beatles["bob"]; !ok {
		fmt.Printf("Bob not found")
	}
}