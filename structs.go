package main

import "fmt"

type User struct {
	Name string
	Email string
}

/**
Main function
 */
func main() {
	u := User{
		Name: "Warren Bailey",
		Email: "warren@warrenbailey.net",
	}
	fmt.Printf("User %v", u.Name)
}
