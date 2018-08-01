package main

import "fmt"

type Users struct {
	ID    int
	First string
	Last  string
}

// Task:
// Satisfy the stringer interface (https://golang.org/pkg/fmt/#Stringer)
// so that the User struct will print
// User <ID> is <First> <Last>
//
// example:
//      User 1 is Rob Pike

func (u Users) String () string {
	return fmt.Sprintf("User %d is %s %s", u.ID, u.First, u.Last)
}

func main() {
	u := Users{ID: 1, First: "Rob", Last: "Pike"}
	fmt.Println(u)
}