package main

/**
 * Write a program that declares two variables
 * 	- one variable to store a VALUE of TYPE int8
 *		- assign to it the largest number possible, then print it
 * 	- one variable to store a VALUE of TYPE uint8
 *		- assign to it the largest number possible, then print it
 * https://go.dev/ref/spec#Numeric_types
 */

import "fmt"

func main() {
	var myInt8 int8 = 127
	fmt.Printf("%v\t%T\n", myInt8, myInt8)
	var myUint8 uint8 = 255
	fmt.Printf("%v\t%T\n", myUint8, myUint8)
}
