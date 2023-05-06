package main

/**
 * Write a program that uses the following:
 * 	- for a VARIABLE storing a VALUE of TYPE
 * 		- string
 * 		- int
 * 		- float64
 * 	- use print verbs to show
 * 		- value
 * 		- type
 */

import "fmt"

func main() {
	myString, myInt, myFloat := "Joh Doe", 42, 3.14
	fmt.Printf("%v\t%T\n", myString, myString)
	fmt.Printf("%v\t%T\n", myInt, myInt)
	fmt.Printf("%v\t%T\n", myFloat, myFloat)
}
