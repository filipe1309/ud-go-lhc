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

	for _, v := range []person{p1, p2} {
		fmt.Printf("First name: %v Last name: %v\n", v.firstName, v.lastName)
		for _, v2 := range v.iceCream {
			fmt.Printf("\t%v\n", v2)
		}
	}
}
