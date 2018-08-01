package main

import "fmt"

func main() {
	names := [4]string{}
	names[0] = "John"
	names[1] = "Paul"
	names[2] = "George"
	names[3] = "Ringo"

	names2 := [4]string{"Warren", "Jimmy", "Bob", "Timmy"}

	names3 := [5]string{}

	fmt.Println(names)
	fmt.Println(names2)
	fmt.Println(names3)

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, n := range names2 {
		fmt.Printf("%d - %s\n", i, n)
	}
}
