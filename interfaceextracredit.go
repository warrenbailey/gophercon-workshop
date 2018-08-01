package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type Persons []Person

// implement sort.Interface on Persons

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var people = Persons{
		{
			Name: "Bill",
			Age:  55,
		},
		{
			Name: "John",
			Age:  92,
		},
		{
			Name: "Megan",
			Age:  4,
		},
	}
	fmt.Printf("unsorted: %+v\n", people)
	sort.Sort(people)
	fmt.Printf("sorted: %+v\n", people)
}