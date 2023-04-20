package main

/*
 * Create a person struct
 * Create a func called “changeMe” with a *person as a parameter
 * Change the value stored at the *person address
 * 	Important
 * 		To dereference a struct, use (*value).field
 * 			p1.first
 * 			(*p1).first
 * 		“As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method),
 * 		x.f is shorthand for (*x).f.”
 * 		https://golang.org/ref/spec#Selectors
 * create a value of type person
 * 	print out the value
 * in func main
 * 	call “changeMe”
 * in func main
 * 	print out the value
 */

import (
	"fmt"
)

type person struct {
	first string
}

func changeMe(p *person) {
	// (*p).first = "James"
	p.first = "James" // it works as well, it's a shorthand
}

func (p person) dontChangeMe() {
	p.first = "Bob" // it doesn't work, it's a copy of the value
}

func (p *person) alsoChangeMe() {
	(*p).first = "Ada"
}

func main() {
	p1 := person{
		first: "John",
	}
	fmt.Println(p1.first)

	changeMe(&p1)
	fmt.Println(p1.first)

	p1.dontChangeMe()
	fmt.Println(p1.first)

	p1.alsoChangeMe()
	fmt.Println(p1.first)
}
