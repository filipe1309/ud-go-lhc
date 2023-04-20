package main

/**
 * 1. Create a func with the identifier foo that
 *    takes in a variadic parameter of type int
 *    pass in a value of type []int into your func (unfurl the []int)
 *    returns the sum of all values of type int passed in
 * 2. Create a func with the identifier bar that
 *    takes in a parameter of type []int
 *    returns the sum of all values of type int passed in
 */

import (
	"fmt"
)

func main() {
	fooBarValues := []int{1, 2, 3, 4, 5}
	fmt.Println(foo(fooBarValues...))
	fmt.Println(bar(fooBarValues))
}

func foo(x ...int) int {
	return sum(x)
}

func bar(x []int) int {
	return sum(x)
}

func sum(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
