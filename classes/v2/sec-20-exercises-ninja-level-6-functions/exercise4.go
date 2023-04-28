package main

/**
 * Create a user defined struct with
 * 	the identifier “person”
 * 	the fields:
 * 		first
 * 		last
 * 		age
 * attach a method to type person with
 * 	the identifier “speak”
 * 	the method should have the person say their name and age
 * create a value of type person
 * call the method from the value of type person
 */

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and I am", p.age, "years old.")
}

func main() {
	person1 := person{
		first: "John",
		last:  "Doe",
		age:   30,
	}

	person1.speak()
}
