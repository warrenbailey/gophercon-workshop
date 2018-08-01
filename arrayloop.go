package main

import "fmt"

func main() {
	fruits := [4]string{"Banana", "Orange", "Pineapple", "Strawberry"}

	// Use a 'classic' for loop  to print out each fruit in the array, and the
	// corresponding index.
	for i :=0; i < len(fruits); i++  {
		fmt.Printf("Index is %d, fruit is %s\n", i, fruits[i])
	}

	// Use the range keyword to loop over the same array of fruits, again printing
	// out the fruit and the corresponding index.
	for n, v := range fruits {
		fmt.Printf("Index is %d, fruit is %s\n", n, v)
	}

	// Using the keyword `continue`, skip every other fruit (even ones)
	for n, v := range fruits {
		if n % 2 == 0 {
			continue
		}
		fmt.Printf("Index is %d, fruit is %s\n", n, v)
	}
}