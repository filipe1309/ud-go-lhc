package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	iceCream  []string
}

func main() {
	p1 := person{
		firstName: "John",
		lastName:  "Doe",
		iceCream: []string{
			"vanilla",
			"chocolate",
			"strawberry",
		},
	}
	p2 := person{
		firstName: "Bob",
		lastName:  "Dylan",
		iceCream: []string{
			"pineapple",
			"mango",
			"banana",
		},
	}

	personMap := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for k, v := range personMap {
		fmt.Printf("Last name: %v\n", k)
		fmt.Printf("\tFirst name: %v\n", v.firstName)
		for _, v2 := range v.iceCream {
			fmt.Printf("\t\t%v\n", v2)
		}
	}
}
